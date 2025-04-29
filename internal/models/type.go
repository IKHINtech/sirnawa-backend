package models

type (
	HouseStatus           string
	RondaAttendanceStatus string
	IplPaymentStatus      string
)

type Role string

func (r Role) ToString() string {
	return string(r)
}

const (
	RoleAdminRT     Role = "admin_rt"
	RoleWakilRT     Role = "wakil_rt"
	RoleKetuaRT     Role = "ketua_rt"
	RoleSekretaris  Role = "sekretaris"
	RoleBendahara   Role = "bendahara"
	RoleKoordinator Role = "koordinator"
	RoleWarga       Role = "warga"
)

const (
	HouseStatusActive   HouseStatus = "aktif"
	HouseStatusInactive HouseStatus = "tidak aktif"
	HouseStatusContract HouseStatus = "kontrak"
)

const (
	RondaAttendanceStatusHadir      RondaAttendanceStatus = "hadir"
	RondaAttendanceStatusTidakHadir RondaAttendanceStatus = "tidak hadir"
)

const (
	IplPaymentStatusUnpaid IplPaymentStatus = "unpaid"
	IplPaymentStatusPaid   IplPaymentStatus = "paid"
)

// CREATE TYPE house_status AS ENUM ('aktif', 'tidak aktif');
// CREATE TYPE ipl_payment_status AS ENUM ('paid', 'unpaid');
// CREATE TYPE ronda_attendance_status AS ENUM ('hadir', 'tidak hadir');
// CREATE TYPE role AS ENUM ('admin_rt', 'wakil_rt', 'sekretaris','bendahara', 'warga');
