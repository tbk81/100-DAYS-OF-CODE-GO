package randoGen

import (
	"math/rand"
	"time"
)

// const symset = "!?@#$%^&*"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func SliceWithRandoset(length int, slice []string) []string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randoset[seededRand.Intn(len(randoset))]
	}
	return string(b)
}

func Rando(length int, slice []string) []string {
	return SliceWithRandoset(length, randoset)
}
