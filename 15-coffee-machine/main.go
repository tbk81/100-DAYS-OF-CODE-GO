package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Machine struct {
	Water  int
	Milk   int
	Coffee int
	Money  float64
}

type Ingredients struct {
	Water  int `json:"water"`
	Milk   int `json:"milk"`
	Coffee int `json:"coffee"`
}

type Drink struct {
	Ingredients Ingredients `json:"ingredients"`
	Cost        float64     `json:"cost"`
}

type MenuData struct {
	Menu map[string]Drink `json:"Menu"`
}

func ResourceCheck(i Ingredients, r Machine) bool {
	if i.Water > r.Water || i.Milk > r.Milk || i.Coffee > r.Coffee {
		return false
	}
	return true
}

func DeductResources(i Ingredients, m *Machine, cost float64) {
	m.Water -= i.Water
	m.Milk -= i.Milk
	m.Coffee -= i.Coffee
	m.Money += cost
}

func Money() { // This is to take funds and check if correct amount has been entered
	return
}

func main() {
	var coffee MenuData
	usrChoice := ""

	byteValue, err := os.ReadFile("menu.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(byteValue, &coffee)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	resources := Machine{
		Water:  1000,
		Milk:   500,
		Coffee: 300,
		Money:  400.0, // Stored in cents as dollars
	}

loop:
	for {
		fmt.Print("What would you like? (espresso/latte/cappuccino): ")
		fmt.Scanln(&usrChoice)

		switch usrChoice {
		case "off":
			break loop
		case "report":
			fmt.Printf("Water: %v mL\nMilk: %v mL\nCoffee: %v g\nMoney: $%.2f\n", resources.Water, resources.Milk, resources.Coffee, resources.Money/100)
		case "espresso", "latte", "cappuccino":
			drink := coffee.Menu[usrChoice]
			if ResourceCheck(drink.Ingredients, resources) {
				DeductResources(drink.Ingredients, &resources, drink.Cost*100)
				fmt.Printf("You chose %s. It costs $%.2f\n", usrChoice, drink.Cost)
				fmt.Println("Preparing your drink... Enjoy!")
			} else {
				fmt.Println("Sorry, not enough resources to make that drink.")
			}
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}
