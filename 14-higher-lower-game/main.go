package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name           string `json:"name"`
	Follower_count int    `json:"follower_count"`
	Description    string `json:"description"`
	Country        string `json:"country"`
}

type PeopleData struct {
	Data []Person `json:"data"`
}

func main() {
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

	for _, person := range people.Data {
		fmt.Printf("%s from %s has %d followers. Description: %s\n",
			person.Name, person.Country, person.Follower_count, person.Description)
	}
}
