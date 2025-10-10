package main

import (
	"fmt"

	"github.com/NoahTek/head-first-go/c14_testing/prose"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithComma(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithComma(phrases))
	phrases = []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithComma(phrases))
}
