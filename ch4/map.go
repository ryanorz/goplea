package ch4

import (
	"unicode/utf8"
	"os"
	"fmt"
	"bufio"
	"unicode"
	"sort"
	"log"
)

func RankByWordCount(wordfreq map[string]int) PairList {
	pl := make(PairList, len(wordfreq))
	i := 0
	for k, v := range wordfreq {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func CharCount(file string) {
	var counts = make(map[string]int)
	var utflen [utf8.UTFMax+1]int
	var letterCount int
	var digitCount int

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("CharCount: %v\n", err)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanRunes)
	for input.Scan() {
		w := input.Text()
		r, size := utf8.DecodeRuneInString(w)
		counts[w]++
		utflen[size]++
		if unicode.IsDigit(r) {
			digitCount++
		} else if unicode.IsLetter(r) {
			letterCount++
		}
	}
	fmt.Printf("rune\tcount\n")
	for _, v := range RankByWordCount(counts) {
		fmt.Printf("%q\t%d\n", v.Key, v.Value)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Println("letterCount:", letterCount)
	fmt.Println("digitCount:", digitCount)
}


func Wordfreq(file string) {
	var wordcouts =  make(map[string]int)

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Wordfreq: %v\n", err)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordcouts[input.Text()]++
	}
	f.Close()
	for _, v := range RankByWordCount(wordcouts) {
		fmt.Println(v.Key, ":", v.Value)
	}
}
