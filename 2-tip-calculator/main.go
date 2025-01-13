package main

import (
	"fmt"
)

func main() {
	var bill, tip, people, total float64

	fmt.Print("Welcome to the tip calculator!\nEnter the total cost of the bill: ")
	fmt.Scanln(&bill)

	fmt.Print("What percentage would you like to tip: ")
	fmt.Scanln(&tip)

	fmt.Print("How many guests to split the bill: ")
	fmt.Scanln(&people)

	total = ((bill * (tip / 100)) + bill) / people
	fmt.Printf("bill: %v\ntip: %v\npeople: %v\ntotal per guest: %.2f\n", bill, tip, people, total)
}
