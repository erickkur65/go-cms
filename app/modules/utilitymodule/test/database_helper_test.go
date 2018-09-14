package test

import (
	"kudotest/app/modules/utilitymodule"
	"testing"
)

func TestSuccessDatabaseConnection(t *testing.T) {
	databaseConfig := utilitymodule.DatabaseConfig{
		Username: "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     "3306",
		DbName:   "kudotest_db"}
	databaseHelper := utilitymodule.DatabaseHelper{DatabaseConfig: databaseConfig}

	if databaseHelper.Connect() == nil {
		t.Error("Connection database fail")
	}
}
