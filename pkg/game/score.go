package game

import (
	"bufio"
	"fmt"
	"os"
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading high scores:", err)
	}
}
