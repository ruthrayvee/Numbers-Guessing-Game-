package main
import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
)
func main (){
	secret := 7
	fmt.Print("Guess the number between 1 and 10:")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	guess, err :=strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return 
	}
	if guess == secret{
		fmt.Println("You guessed it right!🎉")
	} else { 
	fmt.Println ("Wrong guess :( . Try again next time.")
	}
}