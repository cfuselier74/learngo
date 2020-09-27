package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := strings.ToLower(os.Args[1])

	fmt.Println(msg)
}