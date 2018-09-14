package usermodule

import (
	"database/sql"
	"fmt"
)

// UserRepository to handle pengguna table data
type UserRepository struct {
	Database *sql.DB
}

// GetUsers get all users from db
func (userRepository *UserRepository) GetUsers() []Pengguna {
	var users []Pengguna
	var user Pengguna
	sqlQuery := "select id, email, nama_depan, nama_belakang, umur from pengguna"
	rows, err := userRepository.Database.Query(sqlQuery)

	if err != nil {
		fmt.Println("error when fetch users")
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.NamaDepan,
			&user.NamaBelakang,
			&user.Umur)

		if err != nil {
			fmt.Println("error when scan users")
		}

		users = append(users, user)
	}

	return users
}

// GetUserByID get user by given id
func (userRepository *UserRepository) GetUserByID(userID int64) Pengguna {
	var user Pengguna
	sqlQuery := "select id, email, kata_sandi, nama_depan, nama_belakang, umur from pengguna where id = ?"
	row, err := userRepository.Database.Query(sqlQuery, userID)

	if err != nil {
		fmt.Println("error when fetch user by id")
	}

	if row.Next() {
		err = row.Scan(
			&user.ID,
			&user.Email,
			&user.KataSandi,
			&user.NamaDepan,
			&user.NamaBelakang,
			&user.Umur)

		if err != nil {
			fmt.Println("error when scan user by id")
		}
	}

	return user
}

// IsUserEmailExist validate user email is exist or not
func (userRepository *UserRepository) IsUserEmailExist(email string) bool {
	var user Pengguna
	sqlQuery := "select id, email from pengguna where email = ? limit 1"
	row, err := userRepository.Database.Query(sqlQuery, email)

	if err != nil {
		fmt.Println("error when check user email exist")
	}

	if row.Next() {
		err = row.Scan(&user.ID, &user.Email)

		if err != nil {
			return false
		}

		return true
	}

	return false
}

// ValidateUser valdiate user is in database or not
func (userRepository *UserRepository) ValidateUser(user Pengguna) bool {
	var selectedUser Pengguna
	passwordHashHelper := PasswordHashHelper{}

	sqlQuery := "select email, kata_sandi from pengguna where email = ? limit 1"
	row, _ := userRepository.Database.Query(sqlQuery, user.Email)

	if row.Next() {
		err := row.Scan(&selectedUser.Email, &selectedUser.KataSandi)
		if err != nil {
			return false
		}

		if passwordHashHelper.ComparePassword(selectedUser.KataSandi, user.KataSandi) {
			return true
		}
	}

	return false
}

// CreateNewUser create new user to db
func (userRepository *UserRepository) CreateNewUser(user Pengguna) (int64, error) {
	var err error
	sqlQuery := "insert into pengguna(email, kata_sandi, nama_depan, nama_belakang, umur, waktu_dibuat, waktu_diubah) values(?, ?, ?, ?, ?, now(), now())"
	stmt, err := userRepository.Database.Prepare(sqlQuery)

	if err != nil {
		fmt.Println("error when prepare create new user")
	}

	result, _ := stmt.Exec(
		user.Email,
		user.KataSandi,
		user.NamaDepan,
		user.NamaBelakang,
		user.Umur)

	return result.LastInsertId()
}

// UpdateUser update selected user data from db
func (userRepository *UserRepository) UpdateUser(user Pengguna) error {
	var err error
	var sqlQuery string

	if user.KataSandi == "" {
		sqlQuery = `update pengguna set email = ?, nama_depan = ?, nama_belakang = ?, 
			umur = ?, waktu_diubah = now() where id = ?`
	} else {
		sqlQuery = `update pengguna set email = ?, kata_sandi = ?, nama_depan = ?, nama_belakang = ?, 
			umur = ?, waktu_diubah = now() where id = ?`
	}

	stmt, err := userRepository.Database.Prepare(sqlQuery)

	if err != nil {
		fmt.Println("error when prepare update selected user")
	}

	if user.KataSandi == "" {
		_, err = stmt.Exec(
			user.Email,
			user.NamaDepan,
			user.NamaBelakang,
			user.Umur,
			user.ID)
	} else {
		_, err = stmt.Exec(
			user.Email,
			user.KataSandi,
			user.NamaDepan,
			user.NamaBelakang,
			user.Umur,
			user.ID)
	}

	return err
}
