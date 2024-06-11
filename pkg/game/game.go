package game

import (
	"fmt"
	"math/rand"
)

// PlayGame starts the number guessing game with the given range of numbers and tracks the high score.
// The player is also given hints after 5 attempts.
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

		// Give hints after 5 attempts
		if attempts >= 5 {
			if target%2 == 0 {
				fmt.Println("Hint: The number is even.")
			} else {
				fmt.Println("Hint: The number is odd.")
			}
			if target%3 == 0 {
				fmt.Println("Hint: The number is divisible by 3.")
			}
			if isPrime(target) {
				fmt.Println("Hint: The number is a prime number.")
			}
			if target > 50 {
				fmt.Println("Hint: The number is greater than 50.")
			}
		}
	}

	// Save the high score
	saveHighScore(attempts)

	// Display the high scores
	displayHighScores()
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
