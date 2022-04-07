package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	fmt.Println("Hello", name)

	fmt.Println("Enter your age")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSuffix(age, "\n")
	fmt.Println("Hello", name, "you are", age)

	fmt.Println("What is your job")
	job, _ := reader.ReadString('\n')
	job = strings.TrimSuffix(job, "\n")
	fmt.Println("Hello", name, "you are", age, " and you are a", job)
}
