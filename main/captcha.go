package main

import (
	"strconv"
	"strings"
)

func main() {

	// fmt.Printf("Hello \n World!")
}

const (
	// PatternNumOperStr Pattern such as " 1 + one "
	PatternNumOperStr = 1
	// PatternStrOperNum pattern such as " 1 + one "
	PatternStrOperNum = 2
)

const (
	// Plus xxx
	Plus = 1
	// Minus xxx
	Minus = 2
	// Multiply xxx
	Multiply = 3
)

const (
	one = iota
)

// NumWords xxx
var NumWords = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// OperatorSighs xxx
var OperatorSighs = [...]string{"+", "-", "*"}

// Captcha xxx
func Captcha(pattern, operand1, operator, operand2 int) string {

	rOperand1 := strconv.Itoa(operand1)
	rOperand2 := NumWords[operand2-1]
	rOperator := OperatorSighs[operator-1]

	if PatternNumOperStr == pattern {
		return rOperand1 + " " + rOperator + " " + rOperand2
	}

	if PatternStrOperNum == pattern {
		return rOperand2 + " " + rOperator + " " + rOperand1
	}

	return ""

}

// CalCaptcha xxx
func CalCaptcha(captcha string) int {
	captchaArray := strings.Split(captcha, " ")

	wordMap := map[string]int{}
	wordMap["one"] = 1
	wordMap["two"] = 2
	wordMap["three"] = 3
	wordMap["four"] = 4
	wordMap["five"] = 5
	wordMap["six"] = 6
	wordMap["seven"] = 7
	wordMap["eight"] = 8
	wordMap["nine"] = 9

	sighMap := map[string]int{}
	sighMap["+"] = 1
	sighMap["-"] = 2
	sighMap["*"] = 3

	operand1 := 0

	if v, ok := wordMap[captchaArray[0]]; ok {
		operand1 = v
	} else {
		i, _ := strconv.Atoi(captchaArray[0])
		operand1 = i
	}

	operator := 0
	if v, ok := sighMap[captchaArray[1]]; ok {
		operator = v
	}

	operand2 := 0

	if v, ok := wordMap[captchaArray[2]]; ok {
		operand2 = v
	} else {
		i, _ := strconv.Atoi(captchaArray[2])
		operand2 = i
	}

	if Plus == operator {
		return operand1 + operand2
	}

	if Minus == operator {
		return operand1 - operand2
	}

	if Multiply == operator {
		return operand1 * operand2
	}
	return -1
}
