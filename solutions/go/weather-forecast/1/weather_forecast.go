// Package weather provides tools to forecast the current weather condition of cities.
package weather

var (
	// CurrentCondition holds the current weather condition for the city.
	CurrentCondition string
	// CurrentLocation holds the name of the city for which the weather condition is being reported.
	CurrentLocation string
)

// Forecast returns a string reporting the current weather condition for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return "The current weather in " + city + " is " + condition
}
