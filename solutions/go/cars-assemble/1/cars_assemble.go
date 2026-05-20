// Package cars provides functions to calculate the cost and working cars
// produced per min and hour
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const (
        groupedProduceCost = 95000
        individualProduceCost = 10000
    )
    individualCars := carsCount % 10 
    groupedCars := carsCount / 10
    
    return uint(individualCars * individualProduceCost + groupedCars * groupedProduceCost)
}
