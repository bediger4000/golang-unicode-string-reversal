package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	variants := []string{
		"\u0041\u006d\u0065\u0301\u006c\u0069\u0065",
		"\u0041\u006d\u00e9\u006c\u0069\u0065",
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", variants[n])
}
