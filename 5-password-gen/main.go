package main

import (
	"fmt"
	"math/rand"

	"github.com/tbk81/100-DAYS-OF-CODE-GO/5-password-gen/ransym"
)

func numgen(n int) []int {
	return rand.Perm(n)
}

func symgen(n int) string {
	return ransym.Syms(n)
}

func main() {
	var nums, symbols, letters int
	// var letters, symbols string

	fmt.Println("Welcome to the GO password generator\nHow many numbers would you like?")
	fmt.Scanln(&nums)
	fmt.Println("How many symbols would you like?")
	fmt.Scanln(&symbols)
	fmt.Println("How many letters would you like?")
	fmt.Scanln(&letters)

	randoSyms := symgen(symbols)
	// fmt.Println(symbols)

	randoNums := numgen(nums)
	fmt.Printf("Num slice:\n%v\nLen of slice:\n%v\n", randoNums, len(randoNums))
	fmt.Printf("Num sym:\n%v\nLen of sym:\n%v\n", randoSyms, len(randoSyms))
}
