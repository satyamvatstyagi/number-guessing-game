package main

import (
	"fmt"

	"number-guessing-game/pkg/game"
)

// Main function to start the game and play again
func main() {
	var playAgain string

	// Play the game and ask the player if they want to play again
	for {
		var min, max int
		fmt.Print("Enter the minimum value for the range: ")
		fmt.Scan(&min)
		fmt.Print("Enter the maximum value for the range: ")
		fmt.Scan(&max)
		if min >= max {
			fmt.Println("Invalid range. Minimum must be less than maximum.")
			continue
		}

		// Start the game with the given range
		game.PlayGame(min, max)
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scan(&playAgain)

		// Check if the player wants to play again
		if playAgain != "y" {
			break
		}
	}
}
