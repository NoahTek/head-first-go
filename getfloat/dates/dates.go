// Package dates declares the const DaysInWeek and
// has functions to return the number of days in the number of weeks given and vice versa.
package dates

// DaysInWeek is the amount of days in a week; i.e, 7.
const DaysInWeek int = 7

// WeeksToDays returns the number days in the number of weeks in the argument.
func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

// DaysToWeeks returns the float64 weeks of float64 days in the argument.
func DaysToWeeks(days float64) float64 {
	return days / float64(DaysInWeek)
}
