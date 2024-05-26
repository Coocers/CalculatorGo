package main

import (
	"errors"
	"fmt"
	"strconv"
)

func operationValid(operation string) bool {
	var result bool
	switch operation {
	case "+", "-", "*", "/":
		result = true
	default:
		result = false
	}
	return result
}

func convertString(line string) int {
	number, _ := strconv.Atoi(line)
	return number
}

func arabNumberValid(arabicNumber string) bool {
	number := convertString(arabicNumber)
	return number >= 1 && number <= 10
}

func romanNumberValid(romanNumber string) bool {
	var result bool
	switch romanNumber {
	case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
		result = true
	default:
		result = false
	}
	return result
}

func convertRomanArabic(romanNumber string) (int, error) {
	mapping := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	var result int
	i := 0
	for i < len(romanNumber) {
		value, exists := mapping[string(romanNumber[i])]
		if !exists {
			return 0, errors.New("неверное римское число")
		}
		if i+1 < len(romanNumber) && mapping[string(romanNumber[i:i+2])] > 0 {
			value += mapping[string(romanNumber[i:i+2])]
			i++
		}
		result += value
		i++
	}

	return result, nil
}

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func calculate(operand1 int, operand2 int, operation string) int {
	var result int
	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			panic(errors.New("Деление на ноль невозможно"))
		}
		result = operand1 / operand2
	}
	return result
}

func main() {
	var operand1, operand2, operation string

	for {
		fmt.Println("Введите выражение:")
		_, err := fmt.Scanln(&operand1, &operation, &operand2)
		if err != nil {
			panic(errors.New("Ошибка при чтении ввода, 3 значения "))
			return
		}
		break
	}
	if operationValid(operation) {
		if romanNumberValid(operand1) && romanNumberValid(operand2) {
			arabic1, _ := convertRomanArabic(operand1)
			arabic2, _ := convertRomanArabic(operand2)
			arabicResult := calculate(arabic1, arabic2, operation)
			if arabicResult < 1 {
				panic(errors.New("В римской системе нет отрицательных чисел"))
			} else {
				fmt.Println(Roman(arabicResult))
			}
		} else if arabNumberValid(operand1) && arabNumberValid(operand2) {
			arabic1 := convertString(operand1)
			arabic2 := convertString(operand2)
			arabicResult := calculate(arabic1, arabic2, operation)
			fmt.Println(arabicResult)
		} else {
			panic(errors.New("Используются одновременно разные системы счисления"))
		}

	} else {
		panic(errors.New("Недопустимая операция"))
	}
}
