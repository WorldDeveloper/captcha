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
