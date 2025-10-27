package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomInteger() int64 {
	rand.Seed(time.Now().UnixNano())
	// Generate a random integer between 0 (inclusive) and 100 (exclusive)
	return int64(rand.Intn(100))
}
