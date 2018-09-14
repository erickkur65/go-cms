package test

import (
	"kudotest/app/modules/authmodule"
	"kudotest/app/modules/rolemodule"
	"kudotest/app/modules/usermodule"
	"kudotest/app/modules/utilitymodule"
	"testing"
)

func TestSuccessValidateLoginData(t *testing.T) {
	databaseConfig := utilitymodule.DatabaseConfig{
		Username: "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     "3306",
		DbName:   "kudotest_db"}
	databaseHelper := utilitymodule.DatabaseHelper{DatabaseConfig: databaseConfig}
	dbConnection := databaseHelper.Connect()

	groupRepository := rolemodule.GroupRepository{Database: dbConnection}
	userRepository := usermodule.UserRepository{Database: dbConnection}
	authHelper := authmodule.AuthHelper{
		UserRepository:  userRepository,
		GroupRepository: groupRepository}

	user := usermodule.Pengguna{
		Email:     "andi@gmail.com",
		KataSandi: "admin"}

	if authHelper.ValidateLoginData(user) != "Sukses" {
		t.Error("Validate login data fail")
	}
}
