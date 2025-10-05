package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop!")
}

func (r Robot) Walk() {
	fmt.Println("Powering Legs")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))

	toy = Robot("Botco Ambler")
	toy.MakeSound()

	// type assertion
	robot := toy.(Robot) // implicit type for robot; common
	// will panic if the assertion fails at runtime

	// var robot Robot = toy.(Robot)
	// Head First; explicit; Linter throws warning
	// if underlying concrete type of toy is not Robot,
	// program will panic during runtime.

	// var robot = toy.(Robot)
	// implicit; will panic if assertion fails at runtime

	// should use comma-ok idiom — Gemini

	robot.Walk()
}
