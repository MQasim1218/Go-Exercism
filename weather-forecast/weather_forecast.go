// Package weather ...
package weather

// CurrentCondition stores the condition of current weather.
var CurrentCondition string

// CurrentLocation stores the value of current.
var CurrentLocation string

// Forecast takes the city and condition as parameters and returns the Current Weather Condition at the specified location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
