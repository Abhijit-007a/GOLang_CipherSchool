// Package weather provides forcast of a city.
package main

// CurrentCondition stores the current weather condition in string type.
var CurrentCondition string

// CurrentLocation stores the current location in string type.
var CurrentLocation string

// Forecast returns a string representing the current location and the current weather condition out there.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
