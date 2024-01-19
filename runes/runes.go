package runes

import (
	"fmt"
	"reflect"
)
  


var myString = "❗hello"
var stringLength = len(myString)

// fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)
// Output: myString - Length: 8 - Runes: 6
var myRune = '¿'
// fmt.Printf("myRune type: %T\n", myRune)

 
func run() {
 
    // Creating a rune
    rune1 := 'B'
    rune2 := 'g'
    rune3 := '\a'
 
    // Displaying rune and its type
    fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1,
                             rune1, reflect.TypeOf(rune1))
     
    fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2,
                               rune2, reflect.TypeOf(rune2))
     
    fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3, 
                                 reflect.TypeOf(rune3))
 
}