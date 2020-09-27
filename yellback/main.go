package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	ex := strings.Repeat("!", l)

	s := ex + strings.ToUpper(msg) + ex

	fmt.Println(s)

	fmt.Println(len("very") + len("\"cool\""))
}