package main

import (
	"fmt"

	cars "github.com/hebb00/cars"
	pr "github.com/hebb00/print"
	cards "github.com/hebb00/slice"
	str "github.com/hebb00/strings"
	bj "github.com/hebb00/switch"
)

func main() {
	slice := []int{4,5,6,7,8}
	cars.CalculateCost(3)
	fmt.Printf(str.CleanupMessage("*********\n hi*\n******"))
	fmt.Printf(pr.AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
	item := cards.GetItem(slice,3)
	fmt.Printf("item is %d",item)
	fmt.Printf(bj.FirstTurn("ace", "ace", "jack"))

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
