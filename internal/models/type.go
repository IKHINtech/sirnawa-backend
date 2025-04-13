package models

type (
	HouseStatus         string
	RondaActivityStatus string
	IplPaymentStatus    string
)

const (
	HouseStatusActive   HouseStatus = "aktif"
	HouseStatusInactive HouseStatus = "tidak aktif"
)

const (
	RondaActivityStatusHadir      RondaActivityStatus = "hadir"
	RondaActivityStatusTidakHadir RondaActivityStatus = "tidak hadir"
)

const (
	IplPaymentStatusUnpaid IplPaymentStatus = "unpaid"
	IplPaymentStatusPaid   IplPaymentStatus = "paid"
)

// CREATE TYPE house_status AS ENUM ('aktif', 'tidak aktif');
// CREATE TYPE ipl_payment_status AS ENUM ('paid', 'unpaid');
// CREATE TYPE ronda_activity_status AS ENUM ('hadir', 'tidak hadir');
