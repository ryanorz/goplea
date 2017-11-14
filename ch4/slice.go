package ch4

import "unicode"

func ReverseWithArrayPointer(ptr *[32]int) {
	for i, j := 0, 32-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func RotateWithOneLoop(s []int, n int) {
	size := len(s)
	if n >= size {
		return
	}
	start := make([]int, n-1)
	copy(start, s[:n])
	for i := 0; i < size; i++ {
		if i < n {
			s[i] = s[i+n-1]
		} else {
			s[i] = start[i-n]
		}
	}
}

func RemoveDup(s string) string {
	var res string
	var pre rune
	for _, c := range s {
		if pre != c {
			res += string(c)
			pre = c
		}
	}
	return res
}

func RemoveDupString(a []string) []string {
	size := len(a)
	last := 0
	for i := 1; i < size; i++ {
		if a[last] != a[i] {
			a[last+1] = a[i]
			last++
		}
	}
	return a[:last+1]
}

func RemoveDupSpace(a []byte) []byte {
	idx := 0
	preIsSpace := false
	for _, v := range a {
		if !preIsSpace || !unicode.IsSpace(rune(v)) {
			a[idx] = v
			idx++
			preIsSpace = unicode.IsSpace(rune(v))
		}
	}
	return a[:idx]
}

func ReverseRuneString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
