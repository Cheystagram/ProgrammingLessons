package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
// TODO: implement this function!
func main() {

	// Create a bufio reader variable
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	if name == "Cheyenne" {
        fmt.Println("sucks")

	} else if name == "Jack" {
		fmt.Println("is a teacher")
	} else {
		fmt.Println("i dont know you")
	}
	switch name {
	case "Cheyenne":
		fmt.Println("really sucks")
	case "Jack": 
		fmt.Println("is a human")
	default:
		fmt.Println("I dont know you")
	}
	
	
}
	// Print a message asking for the user's name and read in the response

	// Use an if to print out the correct message

	// Use a switch to print out the correct message

