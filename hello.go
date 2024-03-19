package main

import (
	"fmt"
	"strings"
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
type Car struct { 
    battery int 
    batteryDrain int
    speed int
    distance int
}
var car = Car{
    battery:100,
    batteryDrain:2,
    speed:5,
}
// Welcome greets a person by name.
func printCar(car Car) string {
	return fmt.Sprint(car)
}
// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    str := strings.TrimSpace(oldMsg)
   return strings.Trim(str,"* \n*") 
     
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
        speed: speed,
        batteryDrain: batteryDrain,
        battery:100,
        
    }
}
// TODO: define the 'Track' type struct
type Track struct {
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
        distance: distance,
    }
}


// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    bt:= car.battery - car.batteryDrain
    if bt < 0{
        return car
    }
    return  Car{
      battery: bt,
      batteryDrain: car.batteryDrain,
      speed: car.speed,
      distance: car.speed + car.distance,
  }
}
// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	num := car.battery/car.batteryDrain
    max:= car.speed * num
    if track.distance > max{
        return false
    }else{
    return true
    }
}
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

