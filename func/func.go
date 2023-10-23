package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, apt int) int{
    if apt == 0 {
	return len(layers) * 2
    }
    return len(layers) * apt
    
}