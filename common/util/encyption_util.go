package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Encrypt(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Cannot encrypt password", err)
		return "", err
	}
	return string(fromPassword), nil
}

func ValidatePassword(plain string, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain))
	return err == nil
}
