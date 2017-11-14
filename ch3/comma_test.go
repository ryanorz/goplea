package ch3

import "testing"

func TestAddCommaRecursive(t *testing.T) {
	var tests = []struct{
		input string
		want string
	} {
		{"", ""},
		{"0", "0"},
		{"123", "123"},
		{"12345", "12,345"},
	}
	for _, test := range tests {
		if got := AddCommaRecursive(test.input); got != test.want {
			t.Errorf("CommaRecursive(%s) = %s", test.input, got)
		}
	}
}

func TestAddCommaBytesBufferAdvancedVersion(t *testing.T) {
	var tests = []struct{
		input string
		want string
	} {
		{"", ""},
		{"0", "0"},
		{"123", "123"},
		{"12345", "12,345"},
		{"-123", "-123"},
		{"-12345", "-12,345"},
		{"+12345", "+12,345"},
		{"0.1", "0.1"},
		{"0.123", "0.123"},
		{"0.12345", "0.123,45"},
		{"-12345.12345", "-12,345.123,45"},
	}
	for _, test := range tests {
		if got := AddCommaBytesBufferAdvancedVersion(test.input); got != test.want {
			t.Errorf("CommaBytesBufferAdvancedVersion(%s) = %s", test.input, got)
		}
	}
}

func TestSameCharacter(t *testing.T) {
	var tests = []struct{
		s1 string
		s2 string
		want bool
	} {
		{"", "", true},
		{"a", "a", true},
		{"a", "A", false},
		{"abcd", "bdac", true},
		{"ab", "bd", false},
	}
	for _, test := range tests {
		if got := SameCharacter(test.s1, test.s2); got != test.want {
			t.Errorf("SameCharacter(%s, %s) = %v", test.s1, test.s2, got)
		}
	}
}
