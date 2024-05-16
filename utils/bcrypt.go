package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password").Error()
	}
	return string(result)
}
