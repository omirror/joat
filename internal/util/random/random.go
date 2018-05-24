package random

import (
	"crypto/rand"
)

const (
	letters    = `0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz:;.,/?\|-@£#$%^&*`
	lettersHex = `0123456789abcdef`
)

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil
	}
	return b
}

func RandomString(n int) string {
	bytes := RandomBytes(n)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

func RandomHexString(n int) string {
	bytes := RandomBytes(n)
	for i, b := range bytes {
		bytes[i] = lettersHex[b%byte(len(lettersHex))]
	}
	return string(bytes)
}
