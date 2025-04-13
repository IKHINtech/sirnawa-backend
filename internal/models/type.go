package models

type HouseStatus string

type IplPaymentStatus string

const (
	HouseStatusActive   HouseStatus = "aktif"
	HouseStatusInactive HouseStatus = "tidak aktif"
)

const (
	IplPaymentStatusUnpaid IplPaymentStatus = "unpaid"
	IplPaymentStatusPaid   IplPaymentStatus = "paid"
)

// CREATE TYPE house_status AS ENUM ('aktif', 'tidak aktif');
// CREATE TYPE ipl_payment_status AS ENUM ('paid', 'unpaid');
