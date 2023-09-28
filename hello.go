package main

import (
	"fmt"

	greetings "github.com/hebb00/hello"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("hebbs")
    fmt.Println(message)
    defer fmt.Println("world")

	fmt.Println("hello")
    i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) 
	//slice
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}