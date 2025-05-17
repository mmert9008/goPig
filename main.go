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
	// rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano()) // Replaced deprecated rand.Seed function with the currently supported version
	rng := rand.New(source)

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

	maxScore := 50 // Set the maximum score required to win the game

	playerScores := make([]int, numPlayers) // Initialize a slice to store player scores

	fmt.Printf("Get %d points to win!\n", maxScore) // Print win condition

	gameOver := false // Signals the end of the game, used for main game "while" loop

	for !gameOver { // Main game "while" loop
		for playerIndex := 0; playerIndex < numPlayers; playerIndex++ { // Loop through each player
			fmt.Printf("\n--- Player %d's turn ---\n", playerIndex+1)                               // Print turn information for current player
			fmt.Printf("Player %d's current score: %d\n", playerIndex+1, playerScores[playerIndex]) // Print current player's score

			turnScore := 0 // Initialize score for current turn

			for { // Infinite loop for current turn to continuously roll until player chooses to break out of the loop
				fmt.Printf("Player %d, would you like to roll (Y/N): ", playerIndex+1) // Prompt for roll decision

				wannaRoll, _ := reader.ReadString('\n')                   // Read user input for roll decision
				wannaRoll = strings.TrimSpace(strings.ToLower(wannaRoll)) // Convert to lowercase and remove leading/trailing whitespace

				if wannaRoll != "y" && wannaRoll != "" {
					fmt.Printf("\nPlayer %d ended their turn.\n", playerIndex+1) // Print end of turn message
					break                                                        // Break out of infinite loop for current turn
				}

				currentRoll := roll(rng)                                           // Roll the die
				fmt.Printf("Player %d rolled a %d.\n", playerIndex+1, currentRoll) // Print the roll

				if currentRoll == 1 {
					fmt.Printf("Player %d rolled a 1 and ended their turn with 0 points! Git gud, sucka!\n", playerIndex+1) // Print end of turn message
					turnScore = 0                                                                                           // Reset turn score
					break                                                                                                   // Break out of infinite loop for current turn
				} else {
					turnScore += currentRoll // Add the roll to the turn score
					fmt.Printf("Player %d's turn score: %d\n", playerIndex+1, turnScore)
				}
			}

			playerScores[playerIndex] += turnScore                                                  // Add the turn score to the player's total score
			fmt.Printf("\nPlayer %d's total score: %d\n", playerIndex+1, playerScores[playerIndex]) // Print player's total score after their turn

			if playerScores[playerIndex] >= maxScore { // Check win condition
				fmt.Printf("Player %d wins!\n", playerIndex+1) // Print win message
				gameOver = true                                // End game
				break                                          // Break out of main game loop
			}
		}
	} // End of main game loop

	winningScore := 0
	winnerIndex := -1

	for i, score := range playerScores {
		if score > winningScore {
			winningScore = score
			winnerIndex = i
		}
	}

	if winnerIndex != -1 {
		fmt.Printf("\n\nPlayer %d wins with a score of %d!\n", winnerIndex+1, winningScore)
		fmt.Println("\n--------------------")
		fmt.Println("     GAME OVER!     ")
		fmt.Println("--------------------")
	} else {
		fmt.Println("It's a tie!")
	}
}

func roll(rng *rand.Rand) int {
	// Roll a "6-sided die" by generating a random number from 1 to 6
	return rng.Intn(6) + 1
}
