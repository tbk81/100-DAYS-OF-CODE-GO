package main

import (
	"fmt"
	"math/rand"
	// "github.com/tbk81/100-DAYS-OF-CODE-GO/00-tools/clearScreen"
)

func main() {
	var usrChoice string
	var d, guess int
	num := rand.Intn(100)
	fmt.Println("Welcome to the number guessing game!\nChoose a number between 1 and 100.")
	fmt.Println("Pssst, the number is ", num)
	fmt.Print("Choose a difficulty. Type 'e' for easy or 'h' for hard: ")
	fmt.Scanln(&usrChoice)
	if usrChoice == "e" {
		d = 10
		fmt.Printf("You have %v attempts remaining to guess the number.\n", d)
	} else {
		d = 5
		fmt.Printf("You have %v attempts remaining to guess the number.\n", d)
	}
	fmt.Print("Make a guess: ")
	fmt.Scanln(&guess)
	// clearScreen.Clear()
}
