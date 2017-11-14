package ch2

import "testing"

func TestPopCountByTable(t *testing.T) {
	var tests = []struct{
		input uint64
		want int
	} {
		{0x0, 0},
		{0x1, 1},
		{0x2, 1},
		{0x3, 2},
		{0x4, 1},
		{0x5, 2},
	}
	for _, test := range tests {
		if got := PopCountByTable(test.input); got != test.want {
			t.Errorf("PopCountByTable(%v) = %v", test.input, got)
		}
	}
}

func TestPopCountByLoop(t *testing.T) {
	var tests = []struct{
		input uint64
		want int
	} {
		{0x0, 0},
		{0x1, 1},
		{0x2, 1},
		{0x3, 2},
		{0x4, 1},
		{0x5, 2},
	}
	for _, test := range tests {
		if got := PopCountByLoop(test.input); got != test.want {
			t.Errorf("PopCountByTable(%v) = %v", test.input, got)
		}
	}
}

func TestPopCountByBit(t *testing.T) {
	var tests = []struct{
		input uint64
		want int
	} {
		{0x0, 0},
		{0x1, 1},
		{0x2, 1},
		{0x3, 2},
		{0x4, 1},
		{0x5, 2},
	}
	for _, test := range tests {
		if got := PopCountByBit(test.input); got != test.want {
			t.Errorf("PopCountByTable(%v) = %v", test.input, got)
		}
	}
}

func BenchmarkPopCountByTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByTable(uint64(i))
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(uint64(i))
	}
}

func BenchmarkPopCountByBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByBit(uint64(i))
	}
}
