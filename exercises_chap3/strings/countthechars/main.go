package main

import (
	"fmt"
	//"os"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func main() {	
	// When you run it with "İNANÇ", it should return 5 not 7.

	length :=  utf8.RuneCountInString("İNANÇ")
	fmt.Println(length)
}