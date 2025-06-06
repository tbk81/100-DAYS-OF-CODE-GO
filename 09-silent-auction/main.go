package main

import (
	"fmt"
	"os"
	"os/exec"
)

var auctionDict = make(map[string]int)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func addBid(s string, i int) {
	auctionDict[s] = i
}

func checkWinner() { //string {
	var winner string
	var winningBid int
	for k, v := range auctionDict {
		fmt.Println(k, v)
		if v > winningBid {
			winningBid = v
			winner = k
		}
	}
	fmt.Printf("The winner is %v with a bid of $%v\n", winner, winningBid)
}

func main() {
	var name, action string
	var bid int

	fmt.Println("Welcome to the silent auciton!")
Loop:
	for {
		fmt.Println("Do you want to place a bid? Enter Yes or No")
		fmt.Scanln(&action)
		switch {
		case action == "Yes" || action == "yes":
			fmt.Print("Enter your name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter your bid: ")
			fmt.Scanln(&bid)
			addBid(name, bid)
			clearScreen()
			goto Loop
		case action == "No" || action == "no":
			clearScreen()
			checkWinner()
			break Loop
		case action != "Yes" && action != "yes" && action != "No" && action != "no":
			fmt.Printf("%v is not an option, try again\n", action)
			goto Loop
		}
	}
}
