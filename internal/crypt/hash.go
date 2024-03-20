package crypt

import "golang.org/x/crypto/bcrypt"

func GenerateHash(original string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(original), 10)
}

func CheckHash(original, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(original)) == nil
}
