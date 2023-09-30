package main

import "fmt"


func hi() {

	pow := make([]int, 10)//makes an array it length is ten and they are all 0
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
