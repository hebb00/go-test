package main

import (
	"fmt"

	greetings "github.com/hebb00/hello"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("hebbs")
    fmt.Println(message)
    i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) 
}