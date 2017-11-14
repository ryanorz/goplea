package ch4

import (
	"testing"
)

func TestRemoveDup(t *testing.T) {
	tests := []struct{
		input, want string
	} {
		{"", ""},
		{"a", "a"},
		{"abcd", "abcd"},
		{"aa", "a"},
		{"aabb", "ab"},
		{"你好", "你好"},
		{"你你好", "你好"},
	}
	for _, test := range tests {
		if got := RemoveDup(test.input); got != test.want {
			t.Errorf("RemoveDup(%s) = %s", test.input, got)
		}
	}
}
