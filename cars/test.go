package cars

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    	return float64(productionRate) * (successRate / float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    carsPerHour := CalculateWorkingCarsPerHour(productionRate,successRate)
	return int(carsPerHour) / 60    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := uint(carsCount / 10 )
    indivisualCars := uint(carsCount % 10)
    cost := uint(groupsOfTen * 95000 + indivisualCars * 10000)
    fmt.Print(groupsOfTen,indivisualCars, cost )

    return cost
}
