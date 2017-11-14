package ch4

import (
	"github.com/ryanorz/goplea/ch2"
	"flag"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"io"
	"fmt"
	"crypto/sha512"
)

func Sha256Diff(a, b *[32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		sum += ch2.PopCountByTable(uint64(a[i] ^ b[i]))
	}
	return sum
}

func PrintSha()  {
	t := flag.Int("t", 256, "SHA256 SHA384 SHA512")

	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil && err != io.EOF {
		panic(err)
	}

	switch *t {
	case 384:
		fmt.Printf("SHA384: %x\n", sha512.Sum384(bytes))
	case 512:
		fmt.Printf("SHA512: %x\n", sha512.Sum512(bytes))
	default:
		fmt.Printf("SHA256: %x\n", sha256.Sum256(bytes))
	}
}
