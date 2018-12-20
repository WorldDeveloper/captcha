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

// Captcha xxx
func Captcha(pattern, operand1, operator, operand2 int) string {

	rOperand1 := strconv.Itoa(operand1)
	rOperand2 := ""
	rOperator := ""

	if Plus == operator {
		rOperator = "+"
	}

	if Minus == operator {
		rOperator = "-"
	}

	if Multiply == operator {
		rOperator = "*"
	}

	if operand2 == 1 {
		rOperand2 = "one"
	}

	if operand2 == 2 {
		rOperand2 = "two"
	}

	if operand2 == 3 {
		rOperand2 = "three"
	}

	if operand2 == 4 {
		rOperand2 = "four"
	}

	if operand2 == 5 {
		rOperand2 = "five"
	}

	if operand2 == 6 {
		rOperand2 = "six"
	}

	if operand2 == 7 {
		rOperand2 = "seven"
	}

	if operand2 == 8 {
		rOperand2 = "eight"
	}

	if operand2 == 9 {
		rOperand2 = "nine"
	}

	if PatternNumOperStr == pattern {
		return rOperand1 + " " + rOperator + " " + rOperand2
	}

	if PatternStrOperNum == pattern {
		return rOperand2 + " " + rOperator + " " + rOperand1
	}

	return ""

}
