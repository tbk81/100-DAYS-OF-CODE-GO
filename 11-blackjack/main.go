package main

import (
	"fmt"
)

// var deck = make([]string,52)
var deck = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "10", "10", "10", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "10", "10", "10", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "10", "10", "10", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "10", "10", "10"}

func blackjack() []string {
	return deck
}

func main() {
	var usrChoice string
	fmt.Println("Welcome to blackjack!")
Loop:
	for {
		fmt.Print("Do you want to play a game of blackjack? Type 'y' or 'n': ")
		fmt.Scanln(&usrChoice)
		switch usrChoice {
		case "y":
			fmt.Println("start playing game")
			blackjack()
		case "n":
			break Loop
		}
	}
}
