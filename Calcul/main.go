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

func convertRomanArabic(romanNumber string) int {
	var arabicNumber int
	switch romanNumber {
	case "I":
		arabicNumber = 1
	case "II":
		arabicNumber = 2
	case "III":
		arabicNumber = 3
	case "IV":
		arabicNumber = 4
	case "V":
		arabicNumber = 5
	case "VI":
		arabicNumber = 6
	case "VII":
		arabicNumber = 7
	case "VIII":
		arabicNumber = 8
	case "IX":
		arabicNumber = 9
	case "X":
		arabicNumber = 10
	}
	return arabicNumber
}

func convertArabicRoman(arabicNumber int) string {
	var romanNumber string
	switch arabicNumber {
	case 1:
		romanNumber = "I"
	case 2:
		romanNumber = "II"
	case 3:
		romanNumber = "III"
	case 4:
		romanNumber = "IV"
	case 5:
		romanNumber = "V"
	case 6:
		romanNumber = "VI"
	case 7:
		romanNumber = "VII"
	case 8:
		romanNumber = "VIII"
	case 9:
		romanNumber = "IX"
	case 10:
		romanNumber = "X"
	}
	return romanNumber
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
		result = operand1 / operand2
	}
	return result
}

func main() {
	var operand1 string
	var operand2 string
	var operation string
	fmt.Println("Введите выражение:")
	_, _ = fmt.Scan(&operand1, &operation, &operand2)
	if operationValid(operation) {
		if romanNumberValid(operand1) && romanNumberValid(operand2) {
			arabic1 := convertRomanArabic(operand1)
			arabic2 := convertRomanArabic(operand2)
			arabicResult := calculate(arabic1, arabic2, operation)
			if arabicResult < 1 {
				panic(errors.New("В римской системе нет отрицательных чисел"))
			} else {
				fmt.Println(convertArabicRoman(arabicResult))
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
