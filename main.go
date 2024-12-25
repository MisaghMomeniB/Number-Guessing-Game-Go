package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Constants for difficulty levels
const (
	Easy   = 1
	Medium = 2
	Hard   = 3
)

func main() {
	// Seed the random number generator for randomness
	rand.Seed(time.Now().UnixNano())

	// Prompt user to select difficulty level
	difficulty := selectDifficulty()

	// Generate the random number based on the selected difficulty
	numberToGuess := generateRandomNumber(difficulty)

	// Maximum attempts a user can have
	maxAttempts := 10

	// Start the timer to track time taken for guessing
	startTime := time.Now()

	// Track previous guesses to display at the end
	var previousGuesses []int

	// Main game loop (runs for maxAttempts or until the user guesses correctly)
	for attempts := 0; attempts < maxAttempts; attempts++ {
		// Get the user's guess
		userGuess := getUserGuess()

		// Store the user's guess in the previous guesses list
		previousGuesses = append(previousGuesses, userGuess)

		// Check if the guess is correct, too low, or too high
		if userGuess < numberToGuess {
			fmt.Println("Your guess is too low. Try again!")
		} else if userGuess > numberToGuess {
			fmt.Println("Your guess is too high. Try again!")
		} else {
			// Correct guess
			elapsedTime := time.Since(startTime)
			fmt.Printf("Congratulations! You guessed the number %d correctly in %d attempts.\n", numberToGuess, attempts+1)
			fmt.Printf("It took you %v to guess the correct number.\n", elapsedTime)
			fmt.Printf("Your guesses: %v\n", previousGuesses)
			return
		}
	}