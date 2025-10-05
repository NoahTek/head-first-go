package main

import (
	"fmt"
	"log"

	"github.com/NoahTek/head-first-go/calendar"
)

// Unexported variables, struct fields, functions, and methods can still
// be accessed by exported functions and methods in the same package.
func main() {
	date := calendar.Date{}
	err := date.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	event := calendar.Event{}
	err = event.SetTitle("Birthday!")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(1998)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(0o5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}
