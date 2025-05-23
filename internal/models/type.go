package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type (
	HouseStatus           string
	RondaAttendanceStatus string
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

func (i ItemType) ToString() string {
	return string(i)
}

func (v VerificationStatus) ToString() string {
	return string(v)
}

func (p PaymentMethod) ToString() string {
	return string(p)
}

func (h HouseType) ToString() string {
	return string(h)
}

func (i IplBillStatus) ToString() string {
	return string(i)
}

func (h HouseStatus) ToString() string {
	return string(h)
}

type JSONB map[string]any

// Implementasi Scanner dan Valuer untuk JSONB
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &j)
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
	IplBillStatusPaid          IplBillStatus = "paid"           // sudah dibayar
	IplBillStatusUnpaid        IplBillStatus = "unpaid"         // belum dibayar
	IplBillStatusOverdue       IplBillStatus = "overdue"        // telat bayar
	IplBillStatusPenalty       IplBillStatus = "penalty"        // ada denda
	IplBillStatusPartiallyPaid IplBillStatus = "partially_paid" // kurang bayar
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

// CREATE TYPE house_status AS ENUM ('aktif', 'tidak aktif');
// CREATE TYPE ipl_payment_status AS ENUM ('paid', 'unpaid');
// CREATE TYPE ronda_attendance_status AS ENUM ('hadir', 'tidak hadir');
// CREATE TYPE role AS ENUM ('admin_rt', 'wakil_rt', 'sekretaris','bendahara', 'warga');
