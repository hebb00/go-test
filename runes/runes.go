package runes

import "unicode/utf8"

var myString = "❗hello"
var stringLength = len(myString)
var numberOfRunes = utf8.RuneCountInString(myString)

// fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)
// Output: myString - Length: 8 - Runes: 6
var myRune = '¿'
// fmt.Printf("myRune type: %T\n", myRune)
// Output: myRune type: int32


// Output:
// Index: 0	Character: ❗		Code Point: U+2757
// Index: 3	Character: h		Code Point: U+0068
// Index: 4	Character: e		Code Point: U+0065
// Index: 5	Character: l		Code Point: U+006C
// Index: 6	Character: l		Code Point: U+006C
// Index: 7	Character: o		Code Point: U+006F