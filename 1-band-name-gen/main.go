package main

import "fmt"

func main() {
	fmt.Println("Welcome to Band Name Generator.")
	var city, pet string
	fmt.Println("What's the name of the city you grew up in?")
	fmt.Scanf("%c", &city) // This will read user input and assign to var city
	fmt.Println("What's the name of your pet?")
	fmt.Scanf("%p", &pet) // This will read user input and assign to var pet
	fmt.Printf("Your band name could be %c %p.\n")
}
