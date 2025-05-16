package randoGen

// This package generates a random slice of int with n length using any input slice

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func SliceWithRandoset(length int, xi []int) []int {
	b := make([]int, length)
	for i := range b {
		b[i] = xi[seededRand.Intn(len(xi))]
	}
	return b
}

func Rando(length int, xi []int) []int {
	return SliceWithRandoset(length, xi)
}
