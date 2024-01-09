package main

import (
	"fmt"

	cars "github.com/hebb00/cars"
	bird "github.com/hebb00/for"
	f "github.com/hebb00/func"
	pr "github.com/hebb00/print"
	rd "github.com/hebb00/random"
	cards "github.com/hebb00/slice"
	str "github.com/hebb00/strings"
	stc "github.com/hebb00/struct"
	bj "github.com/hebb00/switch"
)



func main() {
	input:=   []float64{0.5, 250, 150, 3, 0.5}
	friendsList:=[]string{"sauce", "noodles", "béchamel", "marjoram"}
	myList:=     []string{"sauce", "noodles", "meat", "tomatoes", "?"}
	slice := []int{4,5,6,7,8}
	cars.CalculateCost(3)
	fmt.Printf(str.CleanupMessage("*********\n hi*\n******"))
	fmt.Printf(pr.AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
	item := cards.GetItem(slice,3)
	fmt.Printf("item is %d",item)
	fmt.Printf(bj.FirstTurn("ace", "ace", "jack"))
	car:=stc.NewCar(5,3)
	fmt.Print(car)
	fmt.Print(bj.ParseCard("jack"))
	fmt.Println(bird.TotalBirdCount(slice))
	fmt.Print(rd.ShuffleAnimals())
	f.AddSecretIngredient(friendsList,myList)
	f.ScaleRecipe(input,5)
}
