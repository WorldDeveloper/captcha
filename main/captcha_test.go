package main_test

import (
	"testing"

	"github.com/abcd/captcha/main"
)

// syn.Waitgroup

func TestCaptchaPaternOneInputONEONEONEONEResult1PlusOne(t *testing.T) {

	expect := "1 + one"
	result := main.Captcha(1, 1, 1, 1)

	if result != expect {
		t.Errorf("expect result : %s, actual : %s ", expect, result)
	}
}

func TestCaptchaPatternOneInputONEONEONETWOResult1PlusTwo(t *testing.T) {
	expect := "1 + two"
	result := main.Captcha(1, 1, 1, 2)

	if result != expect {
		t.Errorf("expect result : %s, actual : %s", expect, result)
	}
}

func TestCaptchaPatternOneInputONEONEONETHREEResult1PlusThree(t *testing.T) {
	expect := "1 + three"
	result := main.Captcha(1, 1, 1, 3)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaPatternOneInputONEONEONEFOUResult1PlusFour(t *testing.T) {
	expect := "1 + four"
	result := main.Captcha(1, 1, 1, 4)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
func TestCaptchaPatternOneInputONEONEONEFIVEResult1PlusFive(t *testing.T) {
	expect := "1 + five"
	result := main.Captcha(1, 1, 1, 5)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaPaternOneInputONEONEONESIXResult1PlusSix(t *testing.T) {
	expect := "1 + six"
	result := main.Captcha(1, 1, 1, 6)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaPaternOneInputONEONEONESEVENResult1PlusSeven(t *testing.T) {
	expect := "1 + seven"
	result := main.Captcha(1, 1, 1, 7)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaPaternOneInputONEONEONEEIGHTResult1PlusEight(t *testing.T) {
	expect := "1 + eight"
	result := main.Captcha(1, 1, 1, 8)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}

func TestCaptchaPaternOneInputONEONEONENINEResult1PlusNine(t *testing.T) {
	expect := "1 + nine"
	result := main.Captcha(1, 1, 1, 9)

	if result != expect {
		t.Errorf("expect result: %s, actual: %s", expect, result)
	}
}
