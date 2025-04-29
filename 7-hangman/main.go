package main

import (
	"fmt"
	"strings"
)

func main() {
	var word []string
	var currentWord string
	var currentGuess string
	wordToGuess := "hang"
	lives := len(wordToGuess)

	// Initializes a slice of _ so index can be referenced and changed
	for range len(wordToGuess) {
		word = append(word, "_")
	}
	// fmt.Println(word)

	// Prints the values of each index of slice createing a seemless "string"
	for _, v := range word {
		fmt.Print(v)
	}

	for {
		fmt.Print("\nGuess a letter: ")
		fmt.Scanln(&currentGuess)
		// fmt.Println(strings.Contains(wordToGuess, currentGuess))
		ind := strings.Index(wordToGuess, currentGuess)
		// fmt.Println("Index:", ind)

		if strings.Contains(wordToGuess, currentGuess) {
			word[ind] = currentGuess
		} else {
			lives -= 1
		}

		fmt.Println("Lives:", lives)
		for _, v := range word {
			currentWord = strings.Join(word, v)
			fmt.Print(currentWord)
			// fmt.Print(v)
		}

		if lives == 0 {
			fmt.Println("\nGame Over!")
			break
		}
	}

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
