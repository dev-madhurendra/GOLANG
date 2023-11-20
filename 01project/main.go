package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func welcomeMessage() {
	fmt.Println("Woohooo, Let's start the game !")
}
func gameInstructions() {
	fmt.Println("Here is the instructions of the game.")
	fmt.Println("1. Select any number between 1 to 100")
	fmt.Println("2. Let's see in how many attempts you will able to guess the number.")
}
func startGame() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Set the range for the random number
	min := 1
	max := 100

	// Generate a random number between min and max
	secretNumber := rand.Intn(max-min+1) + min

	fmt.Printf("I've selected a random number between %d and %d.\n", min, max)

	// Number of allowed guesses
	maxAttempts := 10
	attempts := 0

	for {
		// Get user input
		var guess int
		fmt.Print("What's your guess? ")
		fmt.Scan(&guess)

		// Check the guess
		if guess < min || guess > max {
			fmt.Printf("Please enter a number between %d and %d.\n", min, max)
			continue
		}

		attempts++

		// Provide feedback
		if guess == secretNumber {
			fmt.Printf("Correct! You guessed the number in %d attempts. Well done!\n", attempts)
			break
		} else if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}

		// Check if the maximum number of attempts is reached
		if attempts == maxAttempts {
			fmt.Printf("Sorry, you've reached the maximum number of attempts. The correct number was %d.\n", secretNumber)
			break
		}
	}
}
func playGame() {
	// welcome message
	welcomeMessage()
	// print game instructions
	gameInstructions()
	// game logic
	startGame()
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	//taking input of username
	fmt.Print("Enter you username : ")
	username, _ := reader.ReadString('\n')

	fmt.Println("Welcome, " + username)

	fmt.Println("Wanna play this guessing game ?")
	fmt.Println("Press 1 for YES and 0 for NO")
	// taking input that user wants to play or not
	yesOrNo, _ := reader.ReadString('\n')

	// need to remove the space from string (from front as well as from back)
	removedSpaceYesOrNo := strings.TrimSpace(yesOrNo)

	// If user wants to play then begin the game
	// else exit from the game
	if removedSpaceYesOrNo == "1" {
		playGame()
	} else {
		fmt.Println("OOPs,Sorry might be you are not interested in this game. !!!")
	}
}
