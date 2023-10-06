package main

import (
	"fmt"

	// greetings "github.com/hebb00/hello"
	cars "github.com/hebb00/cars"
)

func main() {
  cars.CalculateCost(3)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
