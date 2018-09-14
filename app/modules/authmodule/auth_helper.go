package authmodule

import (
	"kudotest/app/modules/rolemodule"
	"kudotest/app/modules/usermodule"
)

// AuthHelper handle authentication valdiation
type AuthHelper struct {
	UserRepository  usermodule.UserRepository
	GroupRepository rolemodule.GroupRepository
}

const (
	CustomerRole   = "customer"
	AdminRole      = "admin"
	SuperAdminRole = "super admin"
)

// ValidateLoginData to validate username and password in login page
func (authHelper *AuthHelper) ValidateLoginData(user usermodule.Pengguna) string {
	if user.Email == "" {
		return "Email harus diisi"
	} else if user.KataSandi == "" {
		return "Kata sandi harus diisi"
	} else if !authHelper.UserRepository.ValidateUser(user) {
		return "Email atau kata sandi salah"
	} else if authHelper.GroupRepository.GetUserGroupNameByEmail(user.Email) == CustomerRole {
		// Only admin or super admin role can login to web
		return "Hanya grup admin yang bisa login"
	}

	return "Sukses"
}
