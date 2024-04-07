package random

import (
	"math"
	"math/rand"
)

func GetRandomValue(digits uint8) int32 {
	return rand.Int31n(int32(math.Pow(10, float64(digits))) - 1) + 1
}