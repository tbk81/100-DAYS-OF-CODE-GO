package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Person struct {
	Name          string `json:"name"`
	FollowerCount int    `json:"follower_count"`
	Description   string `json:"description"`
	Country       string `json:"country"`
}

type PeopleData struct {
	Data []Person `json:"data"`
}

func main() {
	var input string
	var score int

	byteValue, err := os.ReadFile("game-data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var people PeopleData
	err = json.Unmarshal(byteValue, &people)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// for _, person := range people.Data {
	// 	fmt.Printf("%s from %s has %d followers. Description: %s\n",
	// 		person.Name, person.Country, person.FollowerCount, person.Description)
	// }

	fmt.Println("Welcome to the higher/lower game!")
	a := people.Data[rand.Intn(len(people.Data))]
	b := people.Data[rand.Intn(len(people.Data))]
Loop:
	for {
		fmt.Printf("\nCompare A: %v, a %v from %v\n", a.Name, a.Description, a.Country)
		fmt.Println("VS")
		fmt.Printf("Compare B: %v, a %v from %v\n", b.Name, b.Description, b.Country)
		fmt.Print("\nWhich account has more followers? Type 'A' or 'B': ")
		fmt.Scanln(&input)
		switch input {
		case "a":
			if a.FollowerCount > b.FollowerCount {
				score += 1
				fmt.Println("You're right! Currnt score:", score)
				b = people.Data[rand.Intn(len(people.Data))]
			} else {
				fmt.Println("That's incorrect, game over. Final score: ", score)
				break Loop
			}
		case "b":
			if b.FollowerCount > a.FollowerCount {
				score += 1
				fmt.Println("You're right! Currnt score:", score)
				a = b
				b = people.Data[rand.Intn(len(people.Data))]
			} else {
				fmt.Println("That's incorrect, game over. Final score: ", score)
				break Loop
			}
		}
	}
}
