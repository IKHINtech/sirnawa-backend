package utils

import (
	"time"
)

func GetWeekRange(t time.Time) (time.Time, time.Time) {
	// Pastikan kita pakai zona waktu lokal (bisa ganti sesuai kebutuhan)
	loc := t.Location()

	// Hitung offset hari ke Senin (Senin = 1, Minggu = 0)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // Jika Minggu, ubah ke 7 supaya bisa dikurangi
	}

	// Hitung start of week (Senin)
	startOfWeek := t.AddDate(0, 0, -weekday+1)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, loc)

	// Hitung end of week (Minggu)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 0, loc)

	return startOfWeek, endOfWeek
}
