package helpers

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 9

	acc := ""
	for i := 0; i < length; i++ {
		acc = acc + strconv.Itoa(rand.Intn(max-min+1)+min)
	}

	return acc
}
