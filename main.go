package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Guess a number between 1 and 10. Enter 0 to quit.")

	randomNumber := rand.Intn(10) + 1
	var guess int

	fmt.Print("Enter your guess: ")
	fmt.Scanln(&guess)

	for guess != 0 && guess != randomNumber {
		fmt.Println("Try again.")
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
	}

	if guess == randomNumber {
		fmt.Printf("Congratulations! You guessed the correct number: %d\n", randomNumber)
	} else {
		fmt.Println("You quit the game.")
	}

	fmt.Println("\nDone.")
}
