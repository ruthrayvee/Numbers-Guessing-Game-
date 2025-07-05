package main 
import (
"fmt"
"math/rand"
"time"
)
func main () {
	rand.Seed(time.Now().UnixNano())
	
	var playAgain string = "yes"
	
	for playAgain == "yes" || playAgain == "Yes" {
	
	var level string 
	maxNumber := 0
	maxTries := 5
	
	fmt.Println ("Welcome to the Numbers Guessing Game!🎉")
	fmt.Println ("Please pick a level : Easy, Medium oe Hard.")
	fmt.Print ("Enter level:")
	fmt.Scan(&level)
	
	switch level {
		case "Easy", "easy", "EASY" :
		maxNumber = 10
		case "Medium", "medium", "MEDIUM" :
		maxNumber = 20
		case "Hard", "hard", "HARD" :
		maxNumber = 40
		default:
		fmt.Println ("Invalid level,defaulting to easy.")
		maxNumber = 10
	}
	
	secretNumber := rand.Intn(maxNumber) + 1
	fmt.Printf("I have picked a number between 1 and %d.You have %d tries!\n", maxNumber, maxTries)
	
	success := false
	
	for i := 1; i <= maxTries; i++ {
		
	var guess int 
	
	fmt.Printf("Attempt %d: Enter your guess: ", i)
	fmt.Scan(&guess)
	
	if guess == secretNumber {
		fmt.Println ("Congratulations!🎆🎆🎆 You guessed it right.!")
		success = true
		break
	} else if guess < secretNumber {
		fmt.Println ("Too low!!")
	} else {
		fmt.Println ("Too high!")
	}
	}
	if !success { 
	fmt.Printf("\n X Game over.The number was: %d\n", secretNumber)
	}
	fmt.Print("\n Do you want to play again? (yes/no):")
	fmt.Scan(&playAgain)
	fmt.Println()
	}
	fmt.Println("Thanks for playing.❤ See you next time!")
}
