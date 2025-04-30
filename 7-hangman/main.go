package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word []string
	var currentWord string
	var currentGuess string
	wordToGuess := "hang"
	lives := len(wordToGuess)

	// Open the file.
	f, _ := os.Open("/home/trevor/go-projects/100-DAYS-GO/7-hangman/hangmanWords.txt")
	defer f.Close()
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

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
		ind := strings.Index(wordToGuess, currentGuess)

		if strings.Contains(wordToGuess, currentGuess) {
			word[ind] = currentGuess
		} else {
			lives -= 1
		}

		fmt.Println("Lives:", lives)
		currentWord = strings.Join(word, "")
		fmt.Println(currentWord)

		// Check word and life count to end the game
		if lives == 0 {
			fmt.Println("\nGame Over!")
			break
		} else if currentWord == wordToGuess {
			fmt.Println("\nYou Win!")
			break
		}
	}

}
