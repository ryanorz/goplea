package ch1

import (
	"os"
	"fmt"
	"bufio"
)

func Dup() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			foundDup := countLines(f, counts)
			f.Close()
			if foundDup {
				fmt.Println(arg)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	foundDup := false
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			foundDup = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	return foundDup
}
