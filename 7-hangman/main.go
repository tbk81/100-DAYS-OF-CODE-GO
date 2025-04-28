package main

import (
	"fmt"
	"strings"
)

func main() {
	var currentGuess string
	word := "hangman"
	for range len(word) {
		fmt.Print("_")
	}
	fmt.Print("\nGuess a letter: ")
	fmt.Scanln(&currentGuess)

	fmt.Println(strings.Contains(word, currentGuess))

}

// for {
// 	var fname, lname string
// 	fmt.Print("Type First Name: ")
// 	fmt.Scanln(&fname)
// 	fmt.Print("Type Last Name: ")
// 	fmt.Scanln(&lname)

// 	fmt.Println("Your name is:", fname, lname)

// 	if fname == "stop" || lname == "stop" {
// 		break
// 	}
// }
