package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func wordGenerator() string {
	var wordList []string
	// Open the file.
	f, _ := os.Open("/home/trevor/go-projects/100-DAYS-GO/7-hangman/hangmanWords.txt")
	defer f.Close()
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and adds them to a list.
	for scanner.Scan() {
		line := scanner.Text()
		wordList = append(wordList, line)
	}
	return wordList[rand.Intn(len(wordList))]
}

// func findAllIndices(word string, char rune) []int {
// 	var indices []int

// 	// Iterate through each character in the word
// 	for i, c := range word {
// 		// If the current character matches the one we're looking for
// 		if c == char {
// 			// Add its index to our results
// 			indices = append(indices, i)
// 		}
// 	}
// 	return indices
// }

// // findAllIndices returns a slice containing all indices
// // where the specified character appears in the word
// func findAllIndices(word string, char string) []int {
// 	var indices []int
// 	// Start searching from the beginning
// 	startIndex := 0

// 	for {
// 		// Find the next occurrence of the character
// 		index := strings.Index(word[startIndex:], char)

// 		// If no more occurrences found, break the loop
// 		if index == -1 {
// 			break
// 		}

// 		// Calculate the actual index in the original string
// 		actualIndex := startIndex + index
// 		indices = append(indices, actualIndex)

// 		// Continue searching after this occurrence
// 		startIndex = actualIndex + 1
// 	}
// 	return indices
// }

func main() {
	var word []string
	var currentWord, currentGuess string
	wordToGuess := wordGenerator()
	lives := 6

	// Initializes a slice of _ so index can be referenced and changed
	for range len(wordToGuess) {
		word = append(word, "_")
	}

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

		// if strings.Contains(wordToGuess, string(currentGuess)) {
		// 	indices := findAllIndices(wordToGuess, currentGuess)
		// 	for _, c := range indices {
		// 		word[c] = string(currentGuess)
		// 	}
		// } else {
		// 	lives -= 1
		// }

		fmt.Println("Lives:", lives)
		currentWord = strings.Join(word, "")
		fmt.Println(currentWord)

		// Check word and life count to end the game
		if lives == 0 {
			fmt.Printf("\nThe word was: %v\nGame Over!\n", wordToGuess)
			break
		} else if currentWord == wordToGuess {
			fmt.Println("\nYou Win!")
			break
		}
	}

}
