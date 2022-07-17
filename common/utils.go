package common

import (
	"math/rand"
)

func GenerateRandomString(strLen uint8, allowedCharSet string) string {
	var charSet []rune

	if len(allowedCharSet) > 0 {
		charSet = []rune(allowedCharSet)
	} else {
		charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	}

	b := make([]rune, strLen)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}
