// average2 calculates the average of serveral numbers.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	simpleCount := float64(len(arguments))
	fmt.Printf("Average: %.2f\n", sum/simpleCount)
}
