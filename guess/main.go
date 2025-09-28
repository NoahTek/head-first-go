// guess challenges players to guess a random number in 10 tries.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess the number?")
	success := false

	for guesses := 10; guesses >= 1; guesses-- {
		fmt.Println("\nYou have", guesses, "guessess left")
		fmt.Print("Make a guess: ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("\nOops. Your guess was low.")
		} else if guess > target {
			fmt.Println("\nOops. Your guess was high.")
		} else {
			success = true
			fmt.Println("\nYou got it! My number is:", target)
			break
		}

	}
	if !success {
		fmt.Print("Sorry, you didn't guess my number. It was: ", target)
	}
}
