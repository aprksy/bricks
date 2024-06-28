package utils

import (
	"crypto/rand"
	"fmt"
)

func RandStr(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid length")
	}
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%x", b)
	return str[:length], nil
}
