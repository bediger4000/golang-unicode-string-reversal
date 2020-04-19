package main

import (
	"fmt"
	"os"

	"golang.org/x/text/unicode/norm"
)

func main() {
	// I think NFC is the right choice for reversing code points

	fmt.Printf("NFC normalized:\n")
	original, reversed := normalizeReverse(norm.NFC, os.Args[1])
	fmt.Printf("%q\n", string(original))
	fmt.Printf("%q\n", string(reversed))

	fmt.Printf("NFD normalized:\n")
	original, reversed = normalizeReverse(norm.NFD, os.Args[1])
	fmt.Printf("%q\n", string(original))
	fmt.Printf("%q\n", string(reversed))
}

func normalizeReverse(form norm.Form, str string) (string, string) {
	normalized := form.String(str)
	original := []rune(normalized)
	reversed := make([]rune, len(original))

	max := len(original)
	l := max - 1

	for i, r := range original {
		reversed[l-i] = r
	}
	return string(original), string(reversed)
}
