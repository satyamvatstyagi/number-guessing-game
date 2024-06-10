package main

import (
	"fmt"

	"number-guessing-game/pkg/game"
)

func main() {
	var playAgain string

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

		game.PlayGame(min, max)
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scan(&playAgain)
		if playAgain != "y" {
			break
		}
	}
}
