// Package weather function to forecast the weather
// and provides variables which stores those information.
package weather



// CurrentCondition represents the weather for a
// given city.
var CurrentCondition string
// CurrentLocation represents the city for the 
// given the weather. 
var	CurrentLocation  string


// Forecast takes city and the condition as argument and gives
// a human readable forecaset.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
