package utilitymodule

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DatabaseHelper helper for handle database
type DatabaseHelper struct {
	DatabaseConfig DatabaseConfig
}

// Connect to connect datatabase
func (databaseHelper *DatabaseHelper) Connect() *sql.DB {
	dsn := databaseHelper.DatabaseConfig.Username + ":" + databaseHelper.DatabaseConfig.Password +
		"@tcp(" + databaseHelper.DatabaseConfig.Host + ":" + databaseHelper.DatabaseConfig.Port +
		")/" + databaseHelper.DatabaseConfig.DbName
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
