package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Band Name Generator.")
	var city, pet string
	fmt.Println("What's the name of the city you grew up in?")
	fmt.Scan(&city) // This will read user input and assign to var city
	fmt.Println("What's the name of your pet?")
	fmt.Scan(&pet) // This will read user input and assign to var pet
	fmt.Printf("Your band name could be %v %v.\n", city, pet)
}

/*
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Band Name Generator!")

	reader := bufio.NewReader(os.Stdin)

	// Get city name
	fmt.Print("What's the name of the city you grew up in? ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	// Get pet name
	fmt.Print("What's the name of your pet? ")
	pet, _ := reader.ReadString('\n')
	pet = strings.TrimSpace(pet)

	// Generate and print band name
	bandName := fmt.Sprintf("%s %s", city, pet)
	fmt.Printf("\nYour band name could be: %s! 🎸\n", bandName)
}
*/
