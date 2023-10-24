package lasagna

import (
	"fmt"
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, apt int) int {
	if apt == 0 {
		return len(layers) * 2
	}
	return len(layers) * apt

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

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(frList, myList []string) {
	i := len(frList) - 1
	x:= len(myList) - 1
	myList[x] = frList[i]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portion int)[]float64 {
	var q = make([]float64,len(amount),cap(amount))
    p:= float64(portion) /2
	for i:=0; i<len(amount); i++{
		q[i]=amount[i]*p
	}
	fmt.Print(amount,q,p)
	return q
}
