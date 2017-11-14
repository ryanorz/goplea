package ch2

import "testing"

func TestCToF(t *testing.T) {
	var tests = [] struct{
		input Celsius
		want Fahrenheit
	} {
		{0, 32},
		{32, 89.6},
		{100, 212},
		{-40, -40},
	}
	for _, test := range tests {
		if got := CToF(test.input); got != test.want {
			t.Errorf("CToF(%v) = %v", test.input, got)
		}
	}
}
