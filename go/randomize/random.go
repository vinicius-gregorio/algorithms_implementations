package randomize

import (
	"math/rand"
)

func GetRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
