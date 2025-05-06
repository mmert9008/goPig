package main

import (
	"bufio"     // For reading user input
	"fmt"       // For printing output
	"math/rand" // For generating random numbers
	"os"        // For reading from standard input
	"strconv"   // For converting input strings to numbers
	"strings"   // For cleaning up input strings
	"time"      // For seeding the random number generator
)

func main() {
	// Seed the random number generator to get a different sequence every time
	rand.Seed(time.Now().UnixNano())

	// "Reader" to read user input
	reader := bufio.NewReader(os.Stdin)

	// Variable for storing the number of players
	var numPlayers int

	// For loop to get the number of players as input
	for {
		fmt.Print("Enter the number of players (2-4): ") // Prompt for the number of players
		input, _ := reader.ReadString('\n')              // Store the user's input if valid, discard error input
		input = strings.TrimSpace(input)                 // Remove leading/trailing whitespace
		n, err := strconv.Atoi(input)                    // Convert the trimmed string to an integer

		// Check for valid input
		if err == nil && n >= 2 && n <= 4 {
			numPlayers = n // Store valid input for the number of players
			break          // Break out of the loop when valid inut is received
		} else {
			fmt.Println("Invalid input. Please enter a number between 2 and 4.") // Print error and continue loop if input is invalid
		}
	}

	fmt.Printf("Welcome to goPig. You've chosen a %d player game.", numPlayers) // Print confirmation with selected number of players
}

func roll() int {
	// Roll a "6-sided die" by generating a random number from 1 to 6
	return rand.Intn(6) + 1
}
