package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}



func mp() {
	var s = map[string]Vertex{
		"Bell Labs":{
			40.68433, -74.39967,
		},
		"Google":{
			37.42202, -122.08408,
		},
	}
	fmt.Println(s)
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // if answer in m ok is true vis the value of the key answer
	fmt.Println("The value:", v, "Present?", ok)
}
