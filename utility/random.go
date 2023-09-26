package utility

import (
	"math/rand"
)

// summon a random number between 0 and 255
func RandomByte() byte {
	return byte(rand.Intn(256))

}
