package ch1

import "testing"

// run "go test -bench=." to get workbench of the two functions.
func BenchmarkForLoopStringsJoin(b *testing.B) {
	testcase := []string{"a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		ForLoopStringsJoin(testcase)
	}
}

func BenchmarkBuiltinStringsJoin(b *testing.B) {
	testcase := []string{"a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		BuiltinStringsJoin(testcase)
	}
}
