package main

import (
	"fmt"
	"os"

	"github.com/gargath/roman/pkg/roman"
)

func main() {
	var input int
	fmt.Print("Please enter a number: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("\n  ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n > %s\n", roman.ToRoman(input))
}
