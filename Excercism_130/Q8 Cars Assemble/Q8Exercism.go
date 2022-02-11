package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var successfulCarPerHour float64 = (successRate / 100) * float64(productionRate)
	return successfulCarPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var successfulCarPerMin int = int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
	return successfulCarPerMin
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	cpforIndivisualCar := 10000
	cpforGroupCar := 95000
	var totalCost uint = ((uint(carsCount) / 10) * uint(cpforGroupCar)) + ((uint(carsCount) % 10) * uint(cpforIndivisualCar))
	return totalCost
}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(37))
}
