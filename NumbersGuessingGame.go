package main

import (
"fmt"
"os"
"bufio"
"strconv"
"strings"
)

func main (){
	secret := 7
	reader := bufio.NewReader(os.Stdin)
	
	maxTries := 5
	tries := 0
	
	fmt.Println ("🎇Welcome to the Numbers Guessing Game!🎆🎆")
	fmt.Println ("Please choose a number between 1-10")
	fmt.Println ("You have", maxTries, "tries.")
	
	for tries < maxTries {
		fmt.Print (" Please enter your number:")
	input , _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	guess,err := strconv.Atoi(input)
	
	
	if err !=nil {
		fmt.Println ("‼Please enter a valid number.")
	continue }
	
	tries++
	
if guess < secret {
fmt.Println ("👇 Too low. Try again. ")
}  else if guess > secret {
fmt.Println	("👆 Too high. Try again. ")
} else if guess == secret { 
fmt.Println ("🎉You guessed it right!!!")
break
}
if tries == maxTries {fmt.Println ("Game over!!!! You've used all your tries. The correct number was",secret)
}
}
}