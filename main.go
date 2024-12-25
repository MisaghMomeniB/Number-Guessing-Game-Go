package main

import (
	"fmt"
	"math/rand"
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

	// Game over if the user runs out of attempts
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nSorry, you've used all %d attempts! The correct number was %d.\n", maxAttempts, numberToGuess)
	fmt.Printf("It took you %v before running out of attempts.\n", elapsedTime)
	fmt.Printf("Your guesses: %v\n", previousGuesses)
}

// selectDifficulty prompts the user to select a difficulty level (Easy, Medium, Hard)
func selectDifficulty() int {
	var difficulty int
	for {
		// Prompt the user for difficulty level selection
		fmt.Println("Select Difficulty Level: 1 (Easy), 2 (Medium), 3 (Hard)")
		fmt.Print("Enter your choice: ")

		// Read the user's input and check for errors
		_, err := fmt.Scanf("%d\n", &difficulty)

		// If there's an error or an invalid choice, prompt the user again
		if err != nil || (difficulty < Easy || difficulty > Hard) {
			fmt.Println("Invalid input! Please enter a number between 1 and 3.")
			// Clear the input buffer to avoid issues with residual data
			clearBuffer()
			continue
		}
		break
	}
	return difficulty
}