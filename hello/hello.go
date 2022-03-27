package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	fmt.Println("Hello", name)
}
