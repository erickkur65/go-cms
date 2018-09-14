package test

import (
	"kudotest/app/modules/usermodule"
	"testing"
)

func TestSuccessValidateCreateNewUser(t *testing.T) {
	userHelper := new(usermodule.UserHelper)
	user := usermodule.Pengguna{
		Email:        "test@yahoo.com",
		KataSandi:    "test",
		NamaDepan:    "ek",
		NamaBelakang: "wijaya"}

	if userHelper.ValidateCreateNewUser(user) != "Sukses" {
		t.Error("Create new user validation fail")
	}
}

func TestSuccessValidateEditProfile(t *testing.T) {
	userHelper := new(usermodule.UserHelper)
	user := usermodule.Pengguna{
		Email:        "test@yahoo.com",
		NamaDepan:    "ek",
		NamaBelakang: "wijaya"}

	if userHelper.ValidateEditProfile(user) != "Sukses" {
		t.Error("Edit profile validation fail")
	}
}
