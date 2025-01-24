package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type calcError string

func (e calcError) Error() string {
	return string(e)
}

func calculate(expression string) error {
	re := regexp.MustCompile(`^\s*\"([^"]{1,10})\"\s*([+\-*/])\s*(\"([^"]{1,10})\"|[1-9]|10)\s*$`)
	// "qwe" + "qwe2"
	// matches[0] "qwe" + "qwe2" 	- вся строка
	// matches[1] qwe 				- первая строка
	// matches[2] + 				- оператор
	// matches[3] "qwe2" 			- вторая строка с кавычками
	// matches[4] qwe2 				- вторая строка (если это строка конечно)
	matches := re.FindStringSubmatch(expression)

	if matches == nil {
		return calcError("Некорректное выражение")
	}

	str1 := matches[1]
	op := matches[2]
	strOrInt := matches[3]

	var result string
	switch op {
	case "+":
		if !strings.HasPrefix(strOrInt, "\"") {
			return calcError("Нельзя складывать строку с числом.")
		}
		str2 := matches[4]
		result = str1 + str2
	case "-":
		if !strings.HasPrefix(strOrInt, "\"") {
			return calcError("Нельзя вычитать число из строки")
		}
		str2 := matches[4]
		result = strings.ReplaceAll(str1, str2, "")
	case "*":
		num, err := strconv.Atoi(strOrInt)
		if err != nil || num < 1 || num > 10 {
			return calcError("Множитель должен быть числом от 1 до 10")
		}
		result = strings.Repeat(str1, num)
	case "/":
		num, err := strconv.Atoi(strOrInt)
		if err != nil || num < 1 || num > 10 {
			return calcError("Делитель должен быть числом от 1 до 10")
		}
		if len(str1) < num {
			result = ""
		} else {
			result = str1[:len(str1)/num]
		}
	default:
		return calcError("Неизвестная операция")
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println("Output:", result)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input:")
	for scanner.Scan() {
		expression := scanner.Text()
		if err := calculate(expression); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			break
		}
		fmt.Println("\nInput:")
	}
}
