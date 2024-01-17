package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Regresa el hash bcrypt del password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// Valida si el password proporcionado es correcto o no
func CheckPassword(password string, hashedPassword string) error {
	fmt.Println(password)
	fmt.Println(hashedPassword)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
