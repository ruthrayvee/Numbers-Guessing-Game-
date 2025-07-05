package main
import ( 
"fmt"
"os"
"bufio"
"strings" )
func main (){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	
	fmt.Print("How are you feeling today, " + name + "?")
	mood, _ := reader.ReadString('\n')
	mood = strings.TrimSpace(mood)
	
	fmt.Println("Hello,", name + "!It's great to know you're feeling", mood, "today.")
}