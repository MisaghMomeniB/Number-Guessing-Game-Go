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