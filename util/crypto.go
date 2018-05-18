package util

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptHashCost = 12
	letters        = `0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz:;.,/?\|-@£#$%^&*`
	lettersHex     = `0123456789abcdef`
)

func EncodeSha1String(value []byte) string {
	h := sha1.New()
	h.Write(value)
	return fmt.Sprintf("%v", hex.EncodeToString(h.Sum(nil)))
}

func EncodeSha256String(value []byte) string {
	h := sha256.New()
	h.Write(value)
	return fmt.Sprintf("%v", hex.EncodeToString(h.Sum(nil)))
}

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptHashCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

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
