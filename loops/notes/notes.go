package main

import (
	"fmt"
)

func infiniteLoop() {
	for {
		fmt.Println("Hello, world!")
	}
}

func main() {
	infiniteLoop()
	// stopEarly()
	// condensedLoop()
	// optionalInitializer()
	// optionalPost()
}

// func stopEarly() {
// 	iterations := 0
// 	for iterations < 5 {
// 		fmt.Println("Hello, world!")
// 		iterations++
// 	}
// 	fmt.Println("Done after", iterations, "iterations.")
// }

// func condensedLoop() {
// 	for iterations := 0; iterations < 5; iterations++ {
// 		fmt.Println("Hello, world!")
// 	}
// 	// fmt.Println("Done after", iterations, "iterations.")
// }

// func optionalInitializer() {
// 	i := 0
// 	for ; i < 5; i++ {
// 		fmt.Println("Hello, world!")
// 	}
// 	fmt.Println("Done after", i, "iterations.")
// }

// func optionalPost() {
// 	for i := 1; i < 5; {
// 		i += i
// 		fmt.Println(i)
// 	}
// }
