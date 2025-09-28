package main

import (
	"fmt"
	"log"

	"github.com/NoahTek/head-first-go/getfloat/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degree celcius.", celcius)
}
