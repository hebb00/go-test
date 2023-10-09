package main

import (
	"fmt"

	// greetings "github.com/hebb00/hello"
	cars "github.com/hebb00/cars"
	str "github.com/hebb00/strings"
)

func main() {
  cars.CalculateCost(3)
  fmt.Printf(str.CleanupMessage("*********\n hi*\n******"))
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
