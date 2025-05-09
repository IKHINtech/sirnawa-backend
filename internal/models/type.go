package models

type (
	HouseStatus           string
	RondaAttendanceStatus string
	IplPaymentStatus      string
	HouseType             string
	ItemType              string
	IplBillStatus         string
	PaymentMethod         string
	VerificationStatus    string
	Role                  string
)

func (r Role) ToString() string {
	return string(r)
}

func (v VerificationStatus) ToString() string {
	return string(v)
}

func (p PaymentMethod) ToString() string {
	return string(p)
}

const (
	Tunai    PaymentMethod = "cash"
	Transfer PaymentMethod = "transfer"
	QRIS     PaymentMethod = "qris"
	EWallet  PaymentMethod = "ewallet"

	Pending  VerificationStatus = "pending"
	Diterima VerificationStatus = "approved"
	Ditolak  VerificationStatus = "rejected"
)

const (
	IplBillStatusPaid          IplBillStatus = "paid"
	IplBillStatusUnpaid        IplBillStatus = "unpaid"
	IplBillStatusOverdue       IplBillStatus = "overdue"
	IplBillStatusPenalty       IplBillStatus = "penalty"
	IplBillStatusPartiallyPaid IplBillStatus = "partially_paid"
)

const (
	ItemTypeIpl ItemType = "ipl"
)

const (
	HouseTtpe60 HouseType = "30/60"
	HouseTtpe72 HouseType = "30/72"
)

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
