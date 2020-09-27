package main

import "fmt"

func main() {
	var (
		msg = "My name is %s and my lastname is %s.\n"
		name = "Chris"
		lastName = "Fuselier" 
	)

	fmt.Printf(msg, name, lastName)
}