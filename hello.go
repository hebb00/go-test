package main

import (
	"fmt"

	cars "github.com/hebb00/cars"
	bird "github.com/hebb00/for"
	f "github.com/hebb00/func"
	m "github.com/hebb00/maps"
	pr "github.com/hebb00/print"
	rd "github.com/hebb00/random"
	cards "github.com/hebb00/slice"
	str "github.com/hebb00/strings"
	stc "github.com/hebb00/struct"
	bj "github.com/hebb00/switch"
	t "github.com/hebb00/time"
)



type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}



func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
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
	// fmt.Print(t.Schedule("2/29/2112 11:59:59"))
	fmt.Print(t.HasPassed("October 3, 2019 20:32:00"))
	fmt.Print(t.Description("7/2/2007 15:04:05"))
	fmt.Print(	t.AnniversaryDate())
	bill:= m.NewBill()
	units:= m.Units()
	m.AddItem(bill, units, "carrot", "dozen")
	m.AddItem(bill, units, "carrot", "half_of_a_dozen")
	m.AddItem(bill, units, "tom", "dozen")
	m.AddItem(bill, units, "tom", "half_of_a_dozen")

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
