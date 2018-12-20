package main

import "strconv"

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

const ()

// Captcha xxx
func Captcha(pattern, operand1, operator, operand2 int) string {
	NumWords := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	OperatorSighs := [...]string{"+", "-", "*"}

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
