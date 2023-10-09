package main

import (
	"fmt"

	// greetings "github.com/hebb00/hello"
	cars "github.com/hebb00/cars"
	comp "github.com/hebb00/comparison"
	str "github.com/hebb00/strings"
)

func main() {
  cars.CalculateCost(3)
  fmt.Printf(str.CleanupMessage("*********\n hi*\n******"))
  fmt.Printf(comp.ChooseVehicle("bmw","ford"))
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
