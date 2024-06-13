# number-guessing-game

This repository contains a console-based Number Guessing Game implemented in Golang. The game randomly selects a number within a range, and the player must guess it. Feedback is provided for each guess (higher/lower), and the game also tracks the number of attempts. It includes enhanced features allowing the player to set the range for the random number, implementing a scoring system, providing hints after a certain number of attempts, and storing high scores in a file to display them.

## Running the Application

There are several ways to run the application:

1. **Using Go Environment:**
   - Install the Go environment.
   - Build and run the application using the command in the number-guessing-game directory: `go run ./main.go`

2. **Using Docker Locally:**
   - Build the Docker image locally using: `docker build -t number-guessing-game-image:1.0.0 .`
   - Run the container using the command: `docker run -tid --name number-guessing-game-container number-guessing-game-image:1.0.0`
   - Exec into the container: check the container ID via `docker ps` and run `docker exec -ti <container_id> sh`
   - Run the application via command: `./number-guessing-game`

3. **Using Docker Image from Docker Hub:**
   - Pull the Docker image from Docker Hub using this command: `docker pull satyamvatstyagi/number-guessing-game-image:1.0.0`
   - Run the container using the command: `docker run -tid --name number-guessing-game-container satyamvatstyagi/number-guessing-game-image:1.0.0`
   - Exec into the container: check the container ID via `docker ps` and run `docker exec -ti <container_id> sh`
   - Run the application via command: `./number-guessing-game`
