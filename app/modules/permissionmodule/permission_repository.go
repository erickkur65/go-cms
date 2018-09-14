package permissionmodule

import (
	"database/sql"
	"fmt"
)

// PermissionRepository to handle akses and grup_akses table data
type PermissionRepository struct {
	Database *sql.DB
}

// GetGroupPermissions
func (permissionRepository *PermissionRepository) GetGroupPermissions(email string) []string {
	var groupPermissions []string
	var groupPermission string

	sqlQuery := `select ak.name from pengguna pa 
		join grup_pengguna gp on pa.id = gp.pengguna_id
		join grup_akses ga on gp.grup_id = ga.grup_id
		join akses ak on ga.akses_id = ak.id where pa.email = ?`
	rows, err := permissionRepository.Database.Query(sqlQuery, email)

	if err != nil {
		fmt.Println("error when fetch group permission")
	}

	for rows.Next() {
		err = rows.Scan(&groupPermission)

		if err != nil {
			fmt.Println("error when scan group permission")
		}

		groupPermissions = append(groupPermissions, groupPermission)
	}

	return groupPermissions
}
