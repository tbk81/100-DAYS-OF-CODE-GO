package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	name           string //`json:"name"`
	follower_count int    //`json:"follower_count"`
	description    string //`json:"description"`
	country        string //`json:"country"`
}

func main() {
	// jFile, err := os.Open("game-data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Sucessfully opened game data")
	// defer jFile.Close()

	byteValue, _ := os.ReadFile("game-data.json")
	var people Person
	json.Unmarshal(byteValue, &people)

	// decoder := json.NewDecoder(jFile)
	// for {
	// 	var output Person
	// 	if err := decoder.Decode(&output); err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("Processed: %+v\n", output)
	// }
}
