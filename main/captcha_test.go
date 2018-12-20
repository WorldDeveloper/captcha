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
