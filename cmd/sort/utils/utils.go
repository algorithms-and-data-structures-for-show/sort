package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Int() - rand.Int()
	}
	return slice
}

func GenerateOrderedSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}
