package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, apt int) int{
    if apt == 0 {
	return len(layers) * 2
    }
    return len(layers) * apt
    
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    var s,n int
    for i:=0; i<len(layers); i++{

        if layers[i] == "noodles"{
            n++
        }
		if layers[i] == "sauce"{
			s++
		}
    }
	return n*50, float64(s)*0.2 
    
}
// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(frList, myList []string){
	i:=len(frList)-1
	myList[i]=frList[i]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portion int)[]float64 {
	q:=amount
	for i:=0; i<len(amount); i++{
		q[i]=amount[i]*float64(portion) 
	}
	return q
}