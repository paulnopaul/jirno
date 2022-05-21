package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return nil, fmt.Errorf("password hashing error: %v", err)
	}
	return hash, err
}

func CmpPassword(password string, hash []byte) bool {
	res := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return res == nil
}
