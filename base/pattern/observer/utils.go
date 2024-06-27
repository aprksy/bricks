package observer

import (
	"crypto/rand"
	"fmt"
)

func randStr(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", b[:length])
}
