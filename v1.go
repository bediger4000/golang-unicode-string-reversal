package main

import (
	"fmt"
	"os"
)

func main() {
	original := []rune(os.Args[1])
	reversed := make([]rune, len(original))
	max := len(original)
	l := max - 1

	for i, r := range original {
		reversed[l-i] = r
	}

	fmt.Printf("%q\n", string(original))
	fmt.Printf("%q\n", string(reversed))
}
