package main

import "fmt"

type (
	Liters      float64
	Milliliters float64
	Gallons     float64
)

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(4)
	fmt.Printf("%v liters equals %.3f Gallons.\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%v milliliters equals %.3f Gallons.\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%v gallons equals %.3f liters.\n", milk, milk.ToLiters())
	fmt.Printf("%v gallons equals %.3f milliliters.\n", milk, milk.ToMilliliters())
}
