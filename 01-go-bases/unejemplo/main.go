package main

import "fmt"

func main() {
	// We got a slice of months and we want to know the season of each month.
	months := []string{"January", "February", "May", "October", "July", "December", "IDK"}

	// We iterate over the slice and for each month we call the GetSeason function.
	for _, selectedMonth := range months {
		season := GetSeason(selectedMonth)
		fmt.Printf("Month %s => Season %s\n", selectedMonth, season)
	}
}

// GetSeason returns the season of the year based on the month.
// If the month is not valid, it returns "Unknown".
func GetSeason(month string) (season string) {
	switch month {
	case "January", "February", "March":
		season = "Summer"
	case "April", "May", "June":
		season = "Autumn"
	case "July", "August", "September":
		season = "Winter"
	case "October", "November", "December":
		season = "Spring"
	default:
		season = "Unknown"
	}
	return
}
