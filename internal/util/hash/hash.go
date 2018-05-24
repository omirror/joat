package hash

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptHashCost = 12
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
