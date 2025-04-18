package ransym

import (
	"math/rand"
	"time"
)

const symset = "!?@#$%^&*"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithSymset(length int, symset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = symset[seededRand.Intn(len(symset))]
	}
	return string(b)
}

func Syms(length int) string {
	return StringWithSymset(length, symset)
}
