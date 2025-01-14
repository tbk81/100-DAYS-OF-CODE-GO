package main

import (
	"fmt"
	"math/rand"
)

func numgen(n int) []int {
	return rand.Perm(n)
}

func main() {
	var nums int
	// var letters, symbols string

	fmt.Println("Welcome to the GO password generator\nHow many numbers would you like?")
	fmt.Scanln(&nums)

	final := numgen(nums)
	fmt.Printf("Num slice:\n%v\nLen of slice:\n%v\n", final, len(final))
}
