package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Чего нибудь посчитать?")

	scanner.Scan()
	text := scanner.Text()
	text = strings.ReplaceAll(text, " ", "")
	arithmeticSign, ex := getArithmeticSign(text)
	if ex != nil {
		fmt.Println(CalculationNotSupportedOperationException{})
		return
	}
	expression := strings.Split(text, arithmeticSign)

	_, err := validation(expression)
	if err != nil {
		fmt.Println(CalculationNotSupportedOperationException{})
		return
	}

	result, e := calculation(expression, arithmeticSign)

	if e != nil {
		fmt.Println(CalculationResultNotSupportedException{})
		return
	}
	fmt.Println(result)
}

func calculation(expression []string, arithmeticSign string) (string, error) {
	isRomanNumber := isRoman(expression[0])

	if isRomanNumber {
		firstNumber, _ := convertToArabic(expression[0])
		secondNumber, _ := convertToArabic(expression[1])
		result, _ := processExpression(firstNumber, secondNumber, arithmeticSign)

		if result < 1 {
			return "", CalculationResultNotSupportedException{}
		} else {
			romanNumerals := make(map[int]string)
			for i := 1; i <= 100; i++ {
				romanNumerals[i] = intToRoman(i)
			}

			romanResult, exists := romanNumerals[result]
			if exists {
				return romanResult, nil
			} else {
				return "", NumberNotFoundException{Message: fmt.Sprintf("%d: данное число не поддерживается калькулятором", result)}
			}
		}
	}

	if !isRomanNumber {
		firstNumber, _ := strconv.Atoi(expression[0])
		secondNumber, _ := strconv.Atoi(expression[1])
		result, _ := processExpression(firstNumber, secondNumber, arithmeticSign)
		return strconv.Itoa(result), nil
	}

	return "", CalculationResultNotSupportedException{}
}

func getArithmeticSign(value string) (string, error) {

	elementsToCheck := []string{"+", "-", "/", "*"}

	for _, element := range elementsToCheck {
		if strings.Contains(value, element) {
			return element, nil
		}
	}

	return "", ArithmeticOperationNotSupportedException{}
}

func validation(value []string) (bool, error) {

	if len(value) > 2 {
		return false, CalculationNotSupportedOperationException{}
	}

	if isArabian(value[0]) && isArabian(value[1]) {
		return true, nil
	}

	if isRoman(value[0]) && isRoman(value[1]) {
		return true, nil
	}

	return false, CalculationNotSupportedOperationException{}
}

func isRoman(value string) bool {

	romanSupportNumbers := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	for _, number := range romanSupportNumbers {
		if number == value {
			return true
		}
	}

	return false
}

func isArabian(value string) bool {

	romanSupportNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for _, number := range romanSupportNumbers {
		if number == value {
			return true
		}
	}

	return false
}

func processExpression(firstValue int, secondValue int, mathOperator string) (int, error) {

	switch mathOperator {
	case "+":
		return firstValue + secondValue, nil
	case "-":
		return firstValue - secondValue, nil
	case "/":
		return firstValue / secondValue, nil
	case "*":
		return firstValue * secondValue, nil
	default:
		return 0, ArithmeticOperationNotSupportedException{}
	}
}

func convertToArabic(value string) (int, error) {
	RomanToArabicMap := make(map[string]int)
	RomanToArabicMap["I"] = 1
	RomanToArabicMap["II"] = 2
	RomanToArabicMap["III"] = 3
	RomanToArabicMap["IV"] = 4
	RomanToArabicMap["V"] = 5
	RomanToArabicMap["VI"] = 6
	RomanToArabicMap["VII"] = 7
	RomanToArabicMap["VIII"] = 8
	RomanToArabicMap["IX"] = 9
	RomanToArabicMap["X"] = 10

	result, exists := RomanToArabicMap[value]
	if exists {
		return result, nil
	} else {
		return 0, NumberNotFoundException{Message: value + ": данное число не поддерживается калькулятором"}
	}
}

func intToRoman(num int) string {
	val := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syb[i]
		}
	}

	return roman
}

type ArithmeticOperationNotSupportedException struct {
	Message string
}

func (e ArithmeticOperationNotSupportedException) Error() string {
	return "Арифмитическая операция не поддерживается"
}

type NumberNotFoundException struct {
	Message string
}

func (e NumberNotFoundException) Error() string {
	return e.Message
}

type CalculationNotSupportedOperationException struct {
	Message string
}

func (e CalculationNotSupportedOperationException) Error() string {
	return "Калькулятор не поддерживает данную операцию"
}

type CalculationResultNotSupportedException struct {
	Message string
}

func (e CalculationResultNotSupportedException) Error() string {
	return "Калькулятор не поддерживает результат арифмитической операции"
}
