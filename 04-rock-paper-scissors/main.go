package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var user, comp int
	comp = rand.Intn(3)

	fmt.Println("Welcome to Rock, Paper, Scissors!\nType 0 for rock, 1 for paper, and 2 for scissors")
	fmt.Scanln(&user)

	switch {
	case user == 0:
		fmt.Println("You chose ROCK")
	case user == 1:
		fmt.Println("You chose PAPER")
	case user == 2:
		fmt.Println("You chose SCISSORS")
	default:
		fmt.Println("You chose invalid number. YOU LOSE!")
		os.Exit(1)
	}

	switch {
	case comp == 0:
		fmt.Println("Computer chose ROCK")
	case comp == 1:
		fmt.Println("Computer chose PAPER")
	default:
		fmt.Println("Computer chose SCISSORS")
	}

	switch {
	case user == comp:
		fmt.Println("It's a draw")
	case user == 0 && comp == 1:
		fmt.Println("You Lose")
	case user == 0 && comp == 2:
		fmt.Println("You Win")
	case user == 1 && comp == 0:
		fmt.Println("You Win")
	case user == 1 && comp == 2:
		fmt.Println("You Lose")
	}

	os.Exit(1)
}
