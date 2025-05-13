package main

import (
	"fmt"
	"os"
)

func main() {
	var choice1, choice2, choice3 string
	fmt.Println("Welcome to treasure island! A choose-your-own adventure game")

	fmt.Print("You come to a fork in the road. Which way do you choose to go, left or right? ")
	fmt.Scanln(&choice1)
	if choice1 == "right" {
		fmt.Print("You run into a bear and it eats you!\nGAME OVER\n")
		os.Exit(1)
	} else {
		fmt.Print("You reach a river that you need to cross. Do you want to wait for a boat or try to swim accross? ")
	}

	fmt.Scanln(&choice2)
	if choice2 == "swim" {
		fmt.Print("You get eaten by and angry beaver!\nGAME OVER\n")
		os.Exit(1)
	} else {
		fmt.Print("A boat comes and take you to the shore.\nYou encounter 3 doors colored red, yellow, and blue. Which door do you choose to go through? ")
	}

	fmt.Scanln(&choice3)
	if choice3 == "red" {
		fmt.Print("You walk into a fireball and die!\nGAME OVER\n")
		os.Exit(1)
	} else if choice3 == "blue" {
		fmt.Print("You walk into a room with a monster and it eats you!\nGAME OVER\n")
		os.Exit(1)
	} else {
		fmt.Print("You found the treasure and now your rich beyond your wildest dreams!\nGAME OVER\n")
	}
}
