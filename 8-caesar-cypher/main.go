package main

import (
	"fmt"
	"slices"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func Encrypt(word string, shift int) string {
	shift = shift % len(alphabet)
	shiftSlice := slices.Concat(alphabet[shift:len(alphabet):len(alphabet)])

}

// function that shifts the indices of a slice
// func translation(a []string, n int) []string {
// 	n = n % len(a)
// 	return slices.Concat(a[n:len(a):len(a)])
// }

func main() {
	// fmt.Println(translation(alphabet, 26))

	var encryptWord, usrChoice string
	var shiftNumber int

	fmt.Println("Type 'endcode' to encrypt, tpye 'decode' to decrypt:")
	fmt.Scanln(&usrChoice)
	if usrChoice == "encode" {
		fmt.Println("Type your message:")
		fmt.Scanln(&encryptWord)
		fmt.Println("Type the shift number:")
		fmt.Scanln(&shiftNumber)
		fmt.Print("Here's your encoded result:", Encrypt(encryptWord))
	} else if usrChoice == "decode" {
		fmt.Println("Type your message:")
		fmt.Scanln(&encryptWord)
	} else {
		fmt.Printf("%v is not an option, try again\n", usrChoice)
	}

}

--------------------------------------------------------------------------------------------------------------------------------------------
package main

import (
	"fmt"
	"slices"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}


func findAllIndices(word string) []int {
	xs := strings.Split(word, "")
	var xi []int
	for _, v := range xs {
		// fmt.Println(alphabet[i])
		for x, _ := range alphabet {
			if v == alphabet[x] {
				xi = append(xi, x)
			}
		}
	}
	return xi
}

func Encrypt(word string, shift int) string {
	encryptSlice
	xs := findAllIndices(word)
	shift = shift % len(alphabet)
	shiftSlice := slices.Concat(alphabet[shift:len(alphabet):len(alphabet)])
	for _,v := range xs{
		for _,y := range shiftSlice{

		}
	}

}

func main() {
	// fmt.Println(translation(alphabet, 26))

	var encryptWord, usrChoice string
	var shiftNumber int

	fmt.Println("Type 'endcode' to encrypt, tpye 'decode' to decrypt:")
	fmt.Scanln(&usrChoice)
	if usrChoice == "encode" {
		fmt.Println("Type your message:")
		fmt.Scanln(&encryptWord)
		fmt.Println("Type the shift number:")
		fmt.Scanln(&shiftNumber)
		fmt.Print("Here's your encoded result:", Encrypt(encryptWord, shiftNumber))
	} else if usrChoice == "decode" {
		fmt.Println("Type your message:")
		fmt.Scanln(&encryptWord)
	} else {
		fmt.Printf("%v is not an option, try again\n", usrChoice)
	}

}


