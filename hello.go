package main

import (
	"fmt"
)

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch  {
	case card =="ace":
		return 11
	case card == "jack" || card =="queen" || card == "king":
		return 10
	case  card == "two":
		return 2
	case  card =="three":
		return 3
	case  card =="four":
		return 4
	case  card == "five":
		return 5
	case  card == "six":
		return 6
	case card == "seven":
		return 7
	case  card == "eight":
		return 8
	case  card == "nine":
		return 9
	case  card == "ten":
		return 10
	default:
		return 0
	}
}
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	score := ParseCard(card1) + ParseCard(card2)
	switch {
	case score == 22:
		return "P"
	case score == 21:
		if ParseCard(dealerCard) != 10 && ParseCard(dealerCard) != 11 {
			return "W"
		} else {
			return "S"
		}
	case score >= 12 && score <= 16:
		if 7 <= ParseCard(dealerCard) {
			return "H"
		} else {
			return "S"
		}
	case score >= 17 && score <= 20:
		return "S"
	case score <= 11:
		return "H"
	default:
		return "H"
	}

}
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var c int
	for k, x := range cb {
		if k == file {
			for _, i := range x {
				if i == true {
					c++
				}

			}
		}

	}
	return c

}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}
// NewTrack creates a new track


func main() {


	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	    type Person struct {
        Name  string
        Likes []string
    }
    var people []*Person

    likes := make(map[string][]*Person)
    for _, p := range people {
        for _, l := range p.Likes {
            likes[l] = append(likes[l], p)
        }
    }
	for _, p := range likes["cheese"] {
        fmt.Println(p.Name, "likes cheese.")
    }
		q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}
// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(frList, myList []string) {
	i := len(frList) - 1
	x:= len(myList) - 1
	myList[x] = frList[i]
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var s, n int
	for i := 0; i < len(layers); i++ {

		if layers[i] == "noodles" {
			n++
		}
		if layers[i] == "sauce" {
			s++
		}
	}
	return n * 50, float64(s) * 0.2

}
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, apt int) int {
	if apt == 0 {
		return len(layers) * 2
	}
	return len(layers) * apt

}

