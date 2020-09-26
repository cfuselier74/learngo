package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	greeterName, age := "gandolf", 2019

	fmt.Println("Hello great", name, "!")
	fmt.Println("My name is", greeterName)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass")
}