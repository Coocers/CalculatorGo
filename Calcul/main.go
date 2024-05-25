package main

import (
	"errors"
	"fmt"
	"strconv"
)

func is_operation_valid(operation string) bool {
	var result bool
	switch operation {
	case "+", "-", "*", "/":
		result = true
	default:
		result = false
	}
	return result
}

func convert_string_to_int(line string) int {
	number, _ := strconv.Atoi(line)
	return number
}

func is_arabic_number_valid(arabic_number string) bool {
	number := convert_string_to_int(arabic_number)
	return number >= 1 && number <= 10
}

func is_roman_number_valid(roman_number string) bool {
	var result bool
	switch roman_number {
	case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
		result = true
	default:
		result = false
	}
	return result
}

func convert_roman_to_arabic(roman_number string) int {
	var arabic_number int
	switch roman_number {
	case "I":
		arabic_number = 1
	case "II":
		arabic_number = 2
	case "III":
		arabic_number = 3
	case "IV":
		arabic_number = 4
	case "V":
		arabic_number = 5
	case "VI":
		arabic_number = 6
	case "VII":
		arabic_number = 7
	case "VIII":
		arabic_number = 8
	case "IX":
		arabic_number = 9
	case "X":
		arabic_number = 10
	}
	return arabic_number
}

func convert_arabic_to_roman(arabic_number int) string {
	var roman_number string
	switch arabic_number {
	case 1:
		roman_number = "I"
	case 2:
		roman_number = "II"
	case 3:
		roman_number = "III"
	case 4:
		roman_number = "IV"
	case 5:
		roman_number = "V"
	case 6:
		roman_number = "VI"
	case 7:
		roman_number = "VII"
	case 8:
		roman_number = "VIII"
	case 9:
		roman_number = "IX"
	case 10:
		roman_number = "X"
	}
	return roman_number
}

func calculate(first_operand int, second_operand int, operation string) int {
	var result int
	switch operation {
	case "+":
		result = first_operand + second_operand
	case "-":
		result = first_operand - second_operand
	case "*":
		result = first_operand * second_operand
	case "/":
		result = first_operand / second_operand
	}
	return result
}

func main() {
	var first_operand string
	var second_operand string
	var operation string
	fmt.Println("Введите выражение:")
	fmt.Scanf("%s %s %s ", &first_operand, &operation, &second_operand)
	if is_operation_valid(operation) {
		if is_roman_number_valid(first_operand) && is_roman_number_valid(second_operand) {
			arabic_number_1 := convert_roman_to_arabic(first_operand)
			arabic_number_2 := convert_roman_to_arabic(second_operand)
			arabic_result := calculate(arabic_number_1, arabic_number_2, operation)
			if arabic_result < 1 {
				panic(errors.New("В римской системе нет отрицательных чисел"))
			} else {
				fmt.Println(convert_arabic_to_roman(arabic_result))
			}
		} else if is_arabic_number_valid(first_operand) && is_arabic_number_valid((second_operand)) {
			arabic_number_1 := convert_string_to_int(first_operand)
			arabic_number_2 := convert_string_to_int(second_operand)
			arabic_result := calculate(arabic_number_1, arabic_number_2, operation)
			fmt.Println((arabic_result))
		} else {
			panic(errors.New("Используются одновременно разные системы счисления"))
		}

	} else {
		panic(errors.New("Недопустимая операция"))
	}
}
