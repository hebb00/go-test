package main

import (
	"fmt"

	// greetings "github.com/hebb00/hello"
	cars "github.com/hebb00/cars"
	pr "github.com/hebb00/print"
	str "github.com/hebb00/strings"
)

func main() {
  cars.CalculateCost(3)
  fmt.Printf(str.CleanupMessage("*********\n hi*\n******"))
  fmt.Printf(pr.AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))

}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
