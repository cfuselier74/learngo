package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	// currently it prints 17
	// it should print 5

	name := "inanç           "
	
	fmt.Println(utf8.RuneCountInString(strings.TrimRight(name, " ")))
}