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
	Cost   float64
}

type Resources struct {
	r map[string]Machine
}

type Ingredients struct {
	Water  int     `json:"water"`
	Milk   int     `json:"milk"`
	Coffee int     `json:"coffee"`
	Cost   float64 `json:"cost"`
}

type Drinks struct {
	Menu map[string]Ingredients `json:"menu"`
}

func main() {

	var resources Resources
	var coffee Drinks
	var usrChoice string

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
	resources.r["Water"] = 1000

loop:
	for {
		fmt.Print("What would you like? (espresso/latte/cappuccino):")
		fmt.Scanln(&usrChoice)

		switch usrChoice {
		case "off":
			break loop
		case "espresso":
			fmt.Println("Your chose espresso")
		case "latte":
			fmt.Println("Your chose latte")
		case "cappuccino":
			fmt.Println("Your chose cappuccino")
		}
	}

	// for name, drink := range coffee.Menu {
	// 	fmt.Printf("%v\t%v%T\t%v\t%v\t%v\n", name, drink.Water, drink.Water, drink.Milk, drink.Coffee, drink.Cost)
	// }

	//fmt.Println("Espresso:", coffee.Menu.Espresso)
	//fmt.Println("Latte:", coffee.Menu.Latte)
	//fmt.Println("Cappuccino:", coffee.Menu.Cappuccino)
}
