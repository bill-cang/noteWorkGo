package utils

import (
	"math/rand"
	"time"
)

func Shuffle(slice []interface{}) {
	if len(slice) <= 0 {
		return
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

//前包含,后不包含.
func GetRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min == 0 {
		return rand.Intn(max)
	}
	return rand.Intn(max-min) + (min)
}

//前包含,后不包含.
func GetRandNoSeed(min, max int) int {
	if min == 0 {
		return rand.Intn(max)
	}
	return rand.Intn(max-min) + (min)
}
