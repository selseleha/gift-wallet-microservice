package utils

import (
	"fmt"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRandomPhoneNumber() string {
	return fmt.Sprintf("09%d", random(100000000, 900000000))
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
