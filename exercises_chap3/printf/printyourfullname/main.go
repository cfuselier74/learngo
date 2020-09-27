package main

import (
	"fmt"
	"os"
)

func main() {
	msg := "Your name is %s and your lastname is %s\n"

	fmt.Printf(msg, os.Args[1], os.Args[2])
}