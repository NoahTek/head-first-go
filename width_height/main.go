package main

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("height or width cannot be negetive.")
	}
	return width * height / 10, nil
	// fmt.Printf("Area : %.2f square units.\n", width*height/10)
}

func errorHandling(amount float64, err error) {
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f liters of paint needed.\n", amount)
	}
}

func main() {
	var total float64
	amount, err := paintNeeded(4.2, 3)
	errorHandling(amount, err)
	total += amount

	amount, err = paintNeeded(5.2, -3.5)
	errorHandling(amount, err)
	total += amount

	fmt.Printf("Total paint needed: %.2f liters.", total)
}
