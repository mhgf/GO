package main

import "testing"

func ShuldSumCorrect(t *testing.T) {
	test := sum(4, 5)
	result := 9

	if result != test {
		t.Error(textFormat(test, result))
	}
}

func ShuldDubCorrect(t *testing.T) {
	test := sub(4, 5)
	result := -1

	if result != test {
		t.Error(textFormat(test, result))
	}
}

func ShuldMultCorrect(t *testing.T) {
	test := mult(4, 5)
	result := 25

	if result != test {
		t.Error(textFormat(test, result))
	}
}

func ShuldDivCorrect(t *testing.T) {
	test := div(10)
	result := 10

	if result != test {
		t.Error(textFormat(test, result))
	}
}
