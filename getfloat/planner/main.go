// Sample program made to practice const declaration in another package.
package main

import (
	"fmt"

	"github.com/NoahTek/head-first-go/getfloat/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days.")
}
