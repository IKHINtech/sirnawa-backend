package utils

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

// Fungsi untuk mendapatkan grup berdasarkan ID
func GetGroupID(groupID string, groups []*models.RondaGroup) (*models.RondaGroup, bool) {
	for _, group := range groups {
		if group.ID == groupID {
			return group, true
		}
	}
	return nil, false
}

// Fungsi untuk mendapatkan grup berikutnya berdasarkan urutan
func GetNextGroup(currentGroupID string, groups []*models.RondaGroup) (*models.RondaGroup, bool) {
	var currentGroupUrutan int
	foundCurrent := false
	for _, group := range groups {
		if group.ID == currentGroupID {
			currentGroupUrutan = int(group.Order)
			foundCurrent = true
			break
		}
	}
	if !foundCurrent {
		return nil, false
	}

	nextUrutan := (currentGroupUrutan % len(groups)) + 1 // Siklus kembali ke 1

	for _, group := range groups {
		if int(group.Order) == nextUrutan {
			return group, true
		}
	}
	return nil, false // Seharusnya tidak terjadi jika urutan dikelola dengan baik
}

func GenerateListSchedule(startDate, endDate time.Time, lastSchedule models.RondaSchedule, groups []*models.RondaGroup) []models.RondaSchedule {
	generatedSchedules := []models.RondaSchedule{}
	currentDate := lastSchedule.Date.AddDate(0, 0, 7)
	currentGroupID := lastSchedule.GroupID

	for !currentDate.After(endDate) {
		if !currentDate.Before(startDate) {
			nextGroup, found := GetNextGroup(currentGroupID, groups)
			if found {
				generatedSchedules = append(generatedSchedules, models.RondaSchedule{
					Date:    currentDate,
					GroupID: nextGroup.ID,
					Group:   *nextGroup,
				})
				currentGroupID = nextGroup.ID // Update ID grup saat ini untuk iterasi berikutnya
			}
		}
		currentDate = currentDate.AddDate(0, 0, 7)
	}
	return generatedSchedules
}
