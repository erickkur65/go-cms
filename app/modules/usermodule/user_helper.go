package usermodule

// UserHelper helper for user controller
type UserHelper struct{}

// ValidateCreateNewUser validate user data in create user
func (userHelper *UserHelper) ValidateCreateNewUser(user Pengguna) string {
	if user.Email == "" {
		return "Email harus diisi"
	} else if user.KataSandi == "" {
		return "Kata sandi harus diisi"
	} else if user.NamaDepan == "" {
		return "Nama depan harus diisi"
	} else if user.NamaBelakang == "" {
		return "Nama belakang harus diisi"
	}

	return "Sukses"
}

// ValidateEditProfile validate user data in edit user
func (userHelper *UserHelper) ValidateEditProfile(user Pengguna) string {
	if user.Email == "" {
		return "Email harus diisi"
	} else if user.NamaDepan == "" {
		return "Nama depan harus diisi"
	} else if user.NamaBelakang == "" {
		return "Nama belakang harus diisi"
	}

	return "Sukses"
}
