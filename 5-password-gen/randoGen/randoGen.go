package randoGen

import (
	"math/rand"
	"time"
)

// const symset = "!?@#$%^&*"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithRandoset(length int, randoset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randoset[seededRand.Intn(len(randoset))]
	}
	return string(b)
}

func Rando(length int, randoset string) string {
	return StringWithRandoset(length, randoset)
}
