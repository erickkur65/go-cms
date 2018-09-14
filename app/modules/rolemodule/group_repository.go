package rolemodule

import (
	"database/sql"
	"fmt"
)

// GroupRepository to handle grup and grup_pengguna table data
type GroupRepository struct {
	Database *sql.DB
}

// GetGroups get all groups
func (groupRepository *GroupRepository) GetGroups() []Grup {
	var groups []Grup
	var group Grup
	sqlQuery := "select id, name from grup"
	rows, err := groupRepository.Database.Query(sqlQuery)

	if err != nil {
		fmt.Println("error when fetch groups")
	}

	for rows.Next() {
		err = rows.Scan(&group.ID, &group.Name)

		if err != nil {
			fmt.Println("error when scan groups")
		}

		groups = append(groups, group)
	}

	return groups
}

// GetUserGroupIDByUserID get user group id by given user id
func (groupRepository *GroupRepository) GetUserGroupIDByUserID(userID int64) int64 {
	var userGroupID int64

	sqlQuery := "select grup_id from grup_pengguna where pengguna_id = ?"
	row, err := groupRepository.Database.Query(sqlQuery, userID)

	if err != nil {
		fmt.Println("error when fetch user group by user id")
	}

	if row.Next() {
		err = row.Scan(&userGroupID)

		if err != nil {
			fmt.Println("error when scan user group by user id")
		}
	}

	return userGroupID
}

// GetGroupByUserID get group by given user id
func (groupRepository *GroupRepository) GetGroupByUserID(userID int64) Grup {
	var group Grup

	sqlQuery := `select g.id, g.name from grup_pengguna gp 
		join grup g on gp.grup_id = g.id where gp.pengguna_id = ?`
	row, err := groupRepository.Database.Query(sqlQuery, userID)

	if err != nil {
		fmt.Println("error when fetch user group by user id")
	}

	if row.Next() {
		err = row.Scan(&group.ID, &group.Name)

		if err != nil {
			fmt.Println("error when scan user group by user id")
		}
	}

	return group
}

// GetUserGroupNameByEmail get user group name by given email
func (groupRepository *GroupRepository) GetUserGroupNameByEmail(email string) string {
	var groupName string

	sqlQuery := `select g.name from pengguna pa join grup_pengguna gp on pa.id = gp.pengguna_id
		join grup g on gp.grup_id = g.id where pa.email = ?`
	row, err := groupRepository.Database.Query(sqlQuery, email)

	if err != nil {
		fmt.Println("error when fetch group name by email")
	}

	if row.Next() {
		err = row.Scan(&groupName)

		if err != nil {
			fmt.Println("error when scan group name by email")
		}
	}

	return groupName
}

// InsertUserGroup insert user group
func (groupRepository *GroupRepository) InsertUserGroup(userID, userGroupID int64) error {
	var err error
	sqlQuery := "insert into grup_pengguna values(?, ?)"
	stmt, err := groupRepository.Database.Prepare(sqlQuery)

	if err != nil {
		fmt.Println("error when prepare insert user group")
	}

	_, err = stmt.Exec(userID, userGroupID)

	return err
}

// UpdateUserGroup update user group
func (groupRepository *GroupRepository) UpdateUserGroup(userID, userGroupID int64) error {
	var err error
	sqlQuery := "update grup_pengguna set grup_id = ? where pengguna_id = ?"
	stmt, err := groupRepository.Database.Prepare(sqlQuery)

	if err != nil {
		fmt.Println("error when prepare insert user group")
	}

	_, err = stmt.Exec(userGroupID, userID)

	return err
}
