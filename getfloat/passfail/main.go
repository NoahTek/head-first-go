package main

import (
	"fmt"
	"log"

	"github.com/NoahTek/head-first-go/getfloat/keyboard"
)

func main() {
	fmt.Print("Enter grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade == 100 {
		status = "passing: perfect."
	} else if grade >= 60 {
		status = "passing."
	} else {
		status = "failing."
	}

	fmt.Println("A grade of", grade, "is", status)
}
