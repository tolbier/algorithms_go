package pivoter

import (
	"math/rand"
	"time"
)

type randomPivoter struct{}

func NewRandomPivoter() Pivoter {
	rand.Seed(time.Now().UnixNano())
	return &randomPivoter{}
}

func (fp *randomPivoter) Pivot(list []int) (pivot int) {
	return rand.Intn(len(list))
}
