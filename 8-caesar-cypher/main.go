package main

import (
	"fmt"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func findAllIndices(word string) []int {
	xs := strings.Split(word, "")
	var xi []int
	for _, v := range xs {
		for x, _ := range alphabet {
			if v == alphabet[x] {
				xi = append(xi, x)
			}
		}
	}
	return xi
}

func Encrypt(word string, shift int) string {
	shift = shift % len(alphabet)
	// shiftSlice := slices.Concat(alphabet[shift:len(alphabet):len(alphabet)])
	return word
}

func main() {
	var encryptWord, usrChoice string
	var shiftNumber int

	fmt.Println("Type 'endcode' to encrypt, tpye 'decode' to decrypt:")
	fmt.Scanln(&usrChoice)
Loop:
	for {
		switch {
		case usrChoice == "encode":
			fmt.Println("Type your message:")
			fmt.Scanln(&encryptWord)
			fmt.Println("Type the shift number:")
			fmt.Scanln(&shiftNumber)
			fmt.Print("Here's your encoded result:", Encrypt(encryptWord, shiftNumber))
		case usrChoice == "decode":
			fmt.Println("Type your message:")
			fmt.Scanln(&encryptWord)
		case usrChoice == "done":
			fmt.Println("Exiting...")
			break Loop
		case usrChoice != "encode" && usrChoice != "decode":
			fmt.Printf("%v is not an option, try again\n", usrChoice)
			break
		}
	}
}
