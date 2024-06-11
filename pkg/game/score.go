package game

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func saveHighScore(attempts int) {
	file, err := os.OpenFile("highscores.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving high score:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n", attempts))
	if err != nil {
		fmt.Println("Error writing high score:", err)
	}
}

func displayHighScores() {
	file, err := os.Open("highscores.txt")
	if err != nil {
		fmt.Println("Error reading high scores:", err)
		return
	}
	defer file.Close()

	fmt.Println("High Scores:")
	scanner := bufio.NewScanner(file)
	highScores := make(map[string]bool)

	// Read each line of the file
	for scanner.Scan() {
		score := scanner.Text()
		highScores[score] = true
	}

	// Sort the high scores map
	// Extract the keys
	keys := make([]string, 0, len(highScores))
	for k := range highScores {
		keys = append(keys, k)
	}

	// Sort the keys
	sort.Strings(keys)

	// Display the high scores
	for _, k := range keys {
		fmt.Printf("%s\n", k)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading high scores:", err)
	}
}
