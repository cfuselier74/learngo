package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	const (
		feetInMeters float64 = 0.3048
		feetInYards          = feetInMeters / 0.9144
	)

	if len(os.Args) < 2 {
		fmt.Println("Give me number")
	} else if feet, err := strconv.Atoi(os.Args[1]); err == nil {
		meters := float64(feet) * feetInMeters
		yards := math.Round(float64(feet) * feetInYards)

		fmt.Printf("%d feet is %.2f meters.\n", feet, meters)
		fmt.Printf("%d feet is %.2f yards.\n", feet, yards)
	} else {
		fmt.Printf("ERROR: cannot convert, %s, Err: %s", os.Args[1], err)
	}

}
