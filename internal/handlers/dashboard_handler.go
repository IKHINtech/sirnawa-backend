package handlers

import (
	"errors"
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get Data DashboardMobile
// @Summary Get Data DashboardMobile
// @Description Get Data DashboardMobile
// @Tags Dashboard
// @Accept json
// @Produce json
// @Security Bearer
// @Param rt_id query string true "RT ID"
// @Param house_id query string true "House ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /dashboard/mobile [get]
func DashboardMobile(c *fiber.Ctx, driveService utils.DriveService) error {
	r := &utils.ResponseHandler{}
	rtID := c.Query("rt_id")
	houseID := c.Query("house_id")
	if rtID == "" {
		return r.BadRequest(c, []string{"id is required"})
	}

	if houseID == "" {
		return r.Ok(c, nil, "house_id is required", nil)
	}

	db := database.DB

	// get ronda jadwal ronda diminggu ini limit 1
	var schedules *models.RondaSchedule

	now := time.Now()
	start, end := utils.GetWeekRange(now)
	err := db.Preload("Rt").Preload("Group").
		Where("rt_id = ? AND (date BETWEEN ? AND ?)", rtID, start, end).
		First(&schedules).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			schedules = nil
		} else {
			return r.BadRequest(c, []string{"error", err.Error()})
		}
	}

	var totalMember int64

	if schedules != nil {
		err = db.Model(&models.RondaGroupMember{}).Where("group_id = ?", &schedules.GroupID).Count(&totalMember).Error
	}

	if err != nil {
		return r.BadRequest(c, []string{"error", err.Error()})
	}
	// get pengumuman limit 1
	var announcement *models.Announcement

	err = db.Preload("User").Preload("User.Resident").Preload("Rt").Where("rt_id = ?", rtID).Order("created_at DESC").First(&announcement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			announcement = nil
		} else {
			return r.BadRequest(c, []string{"error", err.Error()})
		}
	}

	// get bill ipl pada bulan ini dan tahun ini
	var currenBill *models.IplBill

	err = db.Preload("House").Preload("House.Block").Preload("Rt").Where("month = ? AND year = ? AND rt_id = ? AND house_id = ?", now.Month(), now.Year(), rtID, houseID).First(&currenBill).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			currenBill = nil
		} else {
			return r.BadRequest(c, []string{"error", err.Error()})
		}
	}

	resp := response.DashboardMobileResponse{
		RondaSchedule: response.RondaScheduleModelToRondaScheduleResponse(schedules, &totalMember),
		Announcement:  response.AnnouncementModelToAnnouncementResponse(announcement, driveService),
		IplBill:       response.IplBillModelToIplBillResponse(currenBill),
	}
	// get event limit 1
	return r.Ok(c, resp, "success", nil)
}
