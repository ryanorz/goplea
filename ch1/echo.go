// Package ch1 contains answers to exercises of chapter1.
package ch1

import (
	"os"
	"fmt"
)

// Echo : answers to 1.1 1.2
func Echo() {
	for i, arg := range os.Args {
		fmt.Println(i, ":", arg)
	}
}
