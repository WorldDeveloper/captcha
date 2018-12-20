package main

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
	// MINUS xxx
	MINUS = 2
	// MULTIPLY xxx
	MULTIPLY = 3
)

// Captcha xxx
func Captcha(pattern, operand1, operator, operand2 int) string {

	result := ""

	if pattern == 1 && operand1 == 1 && operator == 1 {
		result = "1 + "
	}

	if operand2 == 1 {
		result += "one"
	}

	if operand2 == 2 {
		result += "two"
	}

	if operand2 == 3 {
		result += "three"
	}

	if operand2 == 4 {
		result += "four"
	}

	if operand2 == 5 {
		result += "five"
	}

	if operand2 == 6 {
		result += "six"
	}

	if operand2 == 7 {
		result += "seven"
	}

	if operand2 == 8 {
		result += "eight"
	}

	if operand2 == 9 {
		result += "nine"
	}

	return result
}
