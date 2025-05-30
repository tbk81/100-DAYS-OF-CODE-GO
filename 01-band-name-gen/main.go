package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // makes a scanner to be used to read input

	fmt.Println("Welcome to Band Name Generator.")
	var city, pet string
	fmt.Println("What's the name of the city you grew up in?")
	scanner.Scan()        // scan input to buffer
	city = scanner.Text() // saves current buffer to var city

	fmt.Println("What's the name of your pet?")
	scanner.Scan()
	pet = scanner.Text()

	fmt.Printf("Your band name could be %v %v.\n", city, pet)
}

/* This cose was generated by claude by running above code through it
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
