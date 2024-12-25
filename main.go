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