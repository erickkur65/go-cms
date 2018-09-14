package usermodule

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHashHelper handle user password hashing
type PasswordHashHelper struct {
}

// HashPassword hash user password
func (passwordHashHelper *PasswordHashHelper) HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), err
}

// ComparePassword compare user password hash with password that input from login page
func (passwordHashHelper *PasswordHashHelper) ComparePassword(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err == nil {
		return true
	}

	return false
}
