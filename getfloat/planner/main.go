package main

import (
	"fmt"

	"github.com/NoahTek/head-first-go/getfloat/keyboard"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+keyboard.DaysInWeek, "days.")
}
