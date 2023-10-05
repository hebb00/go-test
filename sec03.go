package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func test() {
	var i I = T{"hello"}
	i.M()
	var g interface{} //An empty interface may hold values of any type
	describe(g)

	g = 42
	describe(g)

	g = "hello"
	describe(g)
	//Type assertions
	var h interface{} = "hello"

	s := h.(string)
	fmt.Println(s)

	s, ok := h.(string)
	fmt.Println(s, ok)

	f, ok := h.(float64)
	fmt.Println(f, ok)

	f = h.(float64) // panic
	fmt.Println(f)
}


func describe(g interface{}) {
	fmt.Printf("(%v, %T)\n", g, g)
}