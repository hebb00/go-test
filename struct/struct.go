package speed

// TODO: define the 'Car' type struct
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
