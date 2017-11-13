package ch1

import "strings"

// ForLoopStringsJoin : 1.3 slow function
func ForLoopStringsJoin(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

// BuiltinStringsJoin : 1.3 fast builtin function
func BuiltinStringsJoin(args []string) string {
	return strings.Join(args, " ")
}
