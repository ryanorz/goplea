package ch3

import (
	"bytes"
	"strings"
)

// AddCommaRecursive xxx 12345 -> 12,345
func AddCommaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return AddCommaRecursive(s[:n-3]) + "," + s[n-3:]
}

// AddCommaBytesBufferAdvancedVersion support +/-, float numbers
func AddCommaBytesBufferAdvancedVersion(s string) string {
	var buf bytes.Buffer
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	integerPart := ""
	floatPart := ""
	if strings.Contains(s, ".") {
		pointIndex := strings.LastIndex(s, ".")
		integerPart = s[:pointIndex]
		floatPart = s[pointIndex+1:]
	} else {
		integerPart = s
	}

	integerPartLen := len(integerPart)
	if integerPartLen <= 3 {
		buf.WriteString(integerPart)
	} else {
		spare := (integerPartLen-1) % 3 + 1
		buf.WriteString(integerPart[:spare])
		for i := 0; i < integerPartLen; i += 3 {
			if spare+i+3 >= integerPartLen {
				buf.WriteString("," + integerPart[spare+i :])
				break
			} else {
				buf.WriteString("," + integerPart[spare+i : spare+i+3])
			}
		}
	}

	floatPartLen := len(floatPart)
	if floatPartLen > 0 {
		if floatPartLen <= 3 {
			buf.WriteString("." + floatPart)
		} else {
			buf.WriteString("." + floatPart[:3])
			for i := 3; i < floatPartLen; i++ {
				if i+3 >= floatPartLen {
					buf.WriteString("," + floatPart[i:])
					break
				} else {
					buf.WriteString("," + floatPart[i : i+3])
				}
			}
		}
	}
	return buf.String()
}

// SameCharacter only care ASCII
func SameCharacter(s1, s2 string) bool {
	var counts [256]int
	for _, c := range s1 {
		counts[c]++
	}
	for _, c := range s2 {
		counts[c]--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
