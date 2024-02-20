package main

import (
	"fmt"
)



func main() {


	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	    type Person struct {
        Name  string
        Likes []string
    }
    var people []*Person

    likes := make(map[string][]*Person)
    for _, p := range people {
        for _, l := range p.Likes {
            likes[l] = append(likes[l], p)
        }
    }
	for _, p := range likes["cheese"] {
        fmt.Println(p.Name, "likes cheese.")
    }
		q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}
// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(frList, myList []string) {
	i := len(frList) - 1
	x:= len(myList) - 1
	myList[x] = frList[i]
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var s, n int
	for i := 0; i < len(layers); i++ {

		if layers[i] == "noodles" {
			n++
		}
		if layers[i] == "sauce" {
			s++
		}
	}
	return n * 50, float64(s) * 0.2

}
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, apt int) int {
	if apt == 0 {
		return len(layers) * 2
	}
	return len(layers) * apt

}
