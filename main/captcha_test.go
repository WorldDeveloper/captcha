package main_test

import (
	"testing"

	"github.com/abcd/captcha/main"
)

// syn.Waitgroup

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal1(t *testing.T) {

	expect := "1 + one"
	result := main.Captcha(1, 1, 1, 1)

	if result != expect {
		t.Errorf("expect result : %s, actual : %s ", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal2(t *testing.T) {
	expect := "1 + two"
	result := main.Captcha(1, 1, 1, 2)

	if result != expect {
		t.Errorf("expect result : %s, actual : %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal3(t *testing.T) {
	expect := "1 + three"
	result := main.Captcha(1, 1, 1, 3)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal4(t *testing.T) {
	expect := "1 + four"
	result := main.Captcha(1, 1, 1, 4)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaCasePatternEqualOneAndOperand1Equal1AndOperatorEqual1AndOperand2Equal5(t *testing.T) {
	expect := "1 + five"
	result := main.Captcha(1, 1, 1, 5)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal6(t *testing.T) {
	expect := "1 + six"
	result := main.Captcha(1, 1, 1, 6)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperandEqual7(t *testing.T) {
	expect := "1 + seven"
	result := main.Captcha(1, 1, 1, 7)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal8(t *testing.T) {
	expect := "1 + eight"
	result := main.Captcha(1, 1, 1, 8)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual1AndOperand2Equal9(t *testing.T) {
	expect := "1 + nine"
	result := main.Captcha(1, 1, 1, 9)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal1(t *testing.T) {
	expect := "1 - one"
	result := main.Captcha(1, 1, 2, 1)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal2(t *testing.T) {
	expect := "1 - two"
	result := main.Captcha(1, 1, 2, 2)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal3(t *testing.T) {
	expect := "1 - three"
	result := main.Captcha(1, 1, 2, 3)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal4(t *testing.T) {
	expect := "1 - four"
	result := main.Captcha(1, 1, 2, 4)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal5(t *testing.T) {
	expect := "1 - five"
	result := main.Captcha(1, 1, 2, 5)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal6(t *testing.T) {
	expect := "1 - six"
	result := main.Captcha(1, 1, 2, 6)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal7(t *testing.T) {
	expect := "1 - seven"
	result := main.Captcha(1, 1, 2, 7)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaCasePatternEqual1AndOperand1Equal1AndOperatorEqual2AndOperand2Equal8(t *testing.T) {
	expect := "1 - eight"
	result := main.Captcha(1, 1, 2, 8)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
