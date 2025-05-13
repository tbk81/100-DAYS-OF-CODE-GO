package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/tbk81/100-DAYS-OF-CODE-GO/5-password-gen/randoGen"
)

const symSet = "!?@#$%^&*"
const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func numgen(n int) string {
	var numStr string
	xi := rand.Perm(n)
	for i := range xi {
		x := strconv.Itoa(i)
		numStr += x
	}
	return numStr
}

func symgen(n int) string {
	return randoGen.Rando(n, symSet)
}

func chargen(n int) string {
	return randoGen.Rando(n, charSet)
}

func main() {
	var nums, symbols, letters int

	fmt.Print("Welcome to the GO password generator\nHow many numbers would you like? ")
	fmt.Scanln(&nums)
	fmt.Print("How many symbols would you like? ")
	fmt.Scanln(&symbols)
	fmt.Print("How many letters would you like? ")
	fmt.Scanln(&letters)

	randoNums := numgen(nums)
	randoSyms := symgen(symbols)
	randoChars := chargen(letters)

	concat := randoNums + randoSyms + randoChars
	shuffleCat := strings.Split(concat, "")
	rand.Shuffle(len(shuffleCat), func(i, j int) { shuffleCat[i], shuffleCat[j] = shuffleCat[j], shuffleCat[i] })

	var final string
	for _, v := range shuffleCat {
		final += v
	}

	// fmt.Printf("Nums:%v\tLen:%v\n", randoNums, len(randoNums))
	// fmt.Printf("Symbols:%v\tLen:%v\n", randoSyms, len(randoSyms))
	// fmt.Printf("Chars:%v\tLen:%v\n", randoChars, len(randoChars))
	// fmt.Printf("Concat:%v\tLen:%v\n", concat, len(concat))
	fmt.Println("Your password is:", final)

}
