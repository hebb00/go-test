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