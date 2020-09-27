package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1

	name1, name2, name3 := os.Args[1], os.Args[2], os.Args[3]

	fmt.Printf("There are %d people!\n", count)
	fmt.Println("Hello great", name1)
	fmt.Println("Hello great", name2)
	fmt.Println("Hello great", name3)
	fmt.Println("Nice to meet you")
}