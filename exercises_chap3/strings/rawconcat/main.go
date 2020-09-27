package main

import (
	"fmt"
	"os"
)

func main() {
	msg := 
`Hi ` + os.Args[1] + `! 
How are you?`

	fmt.Println(msg)
}