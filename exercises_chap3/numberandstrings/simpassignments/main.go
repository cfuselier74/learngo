package main

import "fmt"

func main() {
	width, height := 10, 2

	//width = width + 1
	//width = width + height
	//width = width - 1
	//width = width - height
	//width = width * 20
	//width = width / 25
	//width = width % 5

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
}