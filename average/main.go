// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/NoahTek/head-first-go/datafile"
)

func main() {
	numbers, err := datafile.GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)
}
