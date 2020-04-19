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

	for i, j := 0, l; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = original[j], original[i]
	}

	fmt.Printf("%q\n", string(original))
	fmt.Printf("%q\n", string(reversed))
}
