package main

import "fmt"

type Vert struct {
	Lat, Long float64
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func mp() {
	var s = map[string]Vert{
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

	t, ok := m["Answer"] // if answer in m ok is true vis the value of the key answer
	fmt.Println("The value:", t, "Present?", ok)
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
