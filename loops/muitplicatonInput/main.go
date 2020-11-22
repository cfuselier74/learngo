package main

import (
	"fmt"
	"os"
	"strconv"
)

const errMessage = "Not Valid Input"

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(errMessage)
		return
	}

	i, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		fmt.Println(errMessage)
		return
	}

	var max = int(i)

	fmt.Printf("%5s", "X")

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}

}
