package main

import (
	"fmt"
	"math/rand"
)

func numChecker(g, n int) {
	if g > n {
		fmt.Println("Too high.")
	} else if g < n {
		fmt.Println("Too low.")
	}
}

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
		for d > 0 {
			fmt.Printf("You have %v attempts remaining to guess the number.\n", d)
			fmt.Print("Make a guess: ")
			fmt.Scanln(&guess)
			if guess != num {
				numChecker(guess, num)
				fmt.Println("Try again.")
				d -= 1
				if d == 0 {
					fmt.Println("You ran out of guesses, you lose!")
					break
				}
			} else if guess == num {
				fmt.Println("You win!")
				break
			}
		}
	} else if usrChoice == "h" {
		d = 5
		for d > 0 {
			fmt.Printf("You have %v attempts remaining to guess the number.\n", d)
			fmt.Print("Make a guess: ")
			fmt.Scanln(&guess)
			if guess != num {
				numChecker(guess, num)
				fmt.Println("Try again.")
				d -= 1
				if d == 0 {
					fmt.Println("You ran out of guesses, you lose!")
					break
				}
			} else if guess == num {
				fmt.Println("You win!")
				break
			}
		}
	}
}
