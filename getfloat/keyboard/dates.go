package keyboard

const DaysInWeek int = 7

func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days float64) float64 {
	return days / float64(DaysInWeek)
}
