package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Create a bufio reader variable
	reader := bufio.NewReader(os.Stdin)

	// Print a message asking for the user's name and read in the response
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	// Use an if to print out the correct message
	if name == "Cheyenne" {
		fmt.Println("Wow, I'm the greatest!")
	} else if name == "Jack" {
		fmt.Println("You're the best teacher")
	} else {
		fmt.Println("I don't know you")
	}

	// Use a switch to print out the correct message
	switch name {
	case "Cheyenne":
		fmt.Println("Wow, I'm the greatest!")
	case "Jack":
		fmt.Println("You're the best teacher")
	default:
		fmt.Println("I don't know you")
	}
}
