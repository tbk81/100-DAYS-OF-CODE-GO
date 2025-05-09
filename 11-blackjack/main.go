package main

import "fmt"

var deck = []string{}

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
		case "n":
			break Loop
		}
	}
}
