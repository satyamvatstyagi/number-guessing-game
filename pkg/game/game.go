package game

import (
	"fmt"
	"math/rand"
)

func PlayGame(min, max int) {
	target := rand.Intn(max-min+1) + min
	var guess, attempts int

	for {
		fmt.Printf("Enter your guess (%d-%d): ", min, max)
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			// Clear the invalid input
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if guess < min || guess > max {
			fmt.Printf("Invalid range. Please enter a number between %d and %d.\n", min, max)
			continue
		}

		attempts++
		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}

		// Provide hint after 5 attempts
		if attempts >= 5 {
			if target%2 == 0 {
				fmt.Println("Hint: The number is even.")
			} else {
				fmt.Println("Hint: The number is odd.")
			}
		}
	}

	saveHighScore(attempts)
	displayHighScores()
}
