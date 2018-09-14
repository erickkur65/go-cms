package test

import (
	"kudotest/app/modules/usermodule"
	"testing"
)

func TestSuccessPaswordMatch(t *testing.T) {
	str := "Password Match"
	passwordHashHelper := new(usermodule.PasswordHashHelper)

	hash, err := passwordHashHelper.HashPassword(str)

	if err != nil {
		t.Error(err)
	}

	if !passwordHashHelper.ComparePassword(hash, str) {
		t.Error("Password does not match")
	}
}
