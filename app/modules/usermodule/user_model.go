package usermodule

import (
	"kudotest/app/modules/rolemodule"
)

// Pengguna table contains the information of each pengguna
type Pengguna struct {
	ID           int64
	Email        string
	KataSandi    string
	NamaDepan    string
	NamaBelakang string
	Umur         int
	WaktuDibuat  string
	WaktuDiubah  string
	Grup         rolemodule.Grup
}
