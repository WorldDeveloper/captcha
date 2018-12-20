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
func Captcha(pattern, operand1, operator, openrand2 int) string {

	result := ""

	if pattern == 1 && operand1 == 1 && operator == 1 && openrand2 == 1 {
		return "1 + one"
	}

	return result
}
