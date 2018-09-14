package main

import (
	"database/sql"
	"fmt"
	"kudotest/app/modules/usermodule"
	"kudotest/app/modules/utilitymodule"
)

func createSchema(db *sql.DB, dbName string) {
	_, err := db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	sqlQuery := `CREATE TABLE pengguna (
		id integer primary key auto_increment,
		email varchar(100),
		kata_sandi varchar(1000),
		nama_depan varchar(100),
		nama_belakang varchar(100),
		umur integer,
		waktu_dibuat datetime,
		waktu_diubah datetime
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrate pengguna table")

	sqlQuery = `CREATE TABLE grup (
		id integer primary key auto_increment,
		name varchar(100)
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrate grup table")

	sqlQuery = `CREATE TABLE grup_pengguna (
		pengguna_id integer,
		grup_id integer,
		primary key (pengguna_id, grup_id),
		FOREIGN KEY (pengguna_id) REFERENCES pengguna(id),
		FOREIGN KEY (grup_id) REFERENCES grup(id)
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrate grup_pengguna table")

	sqlQuery = `CREATE TABLE akses (
		id integer primary key auto_increment,
		name varchar(100)
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrate akses table")

	sqlQuery = `CREATE TABLE grup_akses (
		grup_id integer,
		akses_id integer,
		primary key (grup_id, akses_id),
		FOREIGN KEY (grup_id) REFERENCES grup(id),
		FOREIGN KEY (akses_id) REFERENCES akses(id)
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrate grup_akses table")
}

func generateDummyData(db *sql.DB, dbName string) {
	passHashHelper := usermodule.PasswordHashHelper{}
	passHash, err := passHashHelper.HashPassword("admin")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	userTableCommands := make(map[int]string)
	groupTableCommands := make(map[int]string)

	userTableCommands[0] = "insert into pengguna(email, kata_sandi, nama_depan, nama_belakang," +
		"umur, waktu_dibuat, waktu_diubah)" +
		"values('andi@gmail.com', '" + passHash + "', 'andi', 'wijaya', 25, now(), now());"
	userTableCommands[1] = "insert into pengguna(email, kata_sandi, nama_depan, nama_belakang," +
		"umur, waktu_dibuat, waktu_diubah)" +
		"values('tono@gmail.com', '" + passHash + "', 'tono', 'wijaya', 27, now(), now());"
	userTableCommands[2] = "insert into pengguna(email, kata_sandi, nama_depan, nama_belakang," +
		"umur, waktu_dibuat, waktu_diubah)" +
		"values('budi@gmail.com', '" + passHash + "', 'budi', 'wijaya', 29, now(), now());"

	_, err = db.Exec("insert into grup(name) values ('super admin');")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("insert into grup(name) values ('admin');")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("insert into grup(name) values ('customer');")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("insert into akses(name) values ('read');")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("insert into akses(name) values ('create');")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("insert into akses(name) values ('edit');")
	if err != nil {
		panic(err)
	}

	for _, command := range userTableCommands {
		_, err = db.Exec(command)
		if err != nil {
			panic(err)
		}
	}

	groupTableCommands[10] = `insert into grup_pengguna values (1, 1);`
	groupTableCommands[11] = `insert into grup_pengguna values (2, 2);`
	groupTableCommands[12] = `insert into grup_pengguna values (3, 3);`
	groupTableCommands[13] = `insert into grup_akses values (1, 1);`
	groupTableCommands[14] = `insert into grup_akses values (1, 2);`
	groupTableCommands[15] = `insert into grup_akses values (1, 3);`
	groupTableCommands[16] = `insert into grup_akses values (2, 1);`
	groupTableCommands[17] = `insert into grup_akses values (2, 2);`

	for _, command := range groupTableCommands {
		_, err = db.Exec(command)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Success generate dummy data")
}

func main() {
	databaseConfig := utilitymodule.DatabaseConfig{
		Username: "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     "3306",
		DbName:   ""}
	databaseHelper := utilitymodule.DatabaseHelper{DatabaseConfig: databaseConfig}
	db := databaseHelper.Connect()
	dbName := "kudotest_db"

	createSchema(db, dbName)
	generateDummyData(db, dbName)
}
