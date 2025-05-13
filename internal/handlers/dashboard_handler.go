package handlers

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// Get Data DashboardMobile
// @Summary Get Data DashboardMobile
// @Description Get Data DashboardMobile
// @Tags Dashboard
// @Accept json
// @Produce json
// @Security Bearer
// @Param rt_id query string true "RT ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /dashboard/mobile [get]
func DashboardMobile(c *fiber.Ctx, driveService utils.DriveService) error {
	r := &utils.ResponseHandler{}
	rtID := c.Query("rt_id")
	if rtID == "" {
		return r.BadRequest(c, []string{"id is required"})
	}

	// get ronda jadwal ronda diminggu ini limit 1
	var schedules *models.RondaSchedule
	db := database.DB

	now := time.Now()
	start, end := utils.GetWeekRange(now)
	err := db.Preload("Rt").Preload("Group").
		Where("rt_id = ? AND (date BETWEEN ? AND ?)", rtID, start, end).
		First(&schedules).Error
	if err != nil {
		return r.BadRequest(c, []string{"error", err.Error()})
	}

	// get pengumuman limit 1
	var announcement *models.Announcement

	err = db.Preload("User").Preload("User.Resident").Preload("Rt").Where("rt_id = ?", rtID).Order("created_at DESC").First(&announcement).Error
	if err != nil {
		return r.BadRequest(c, []string{"error", err.Error()})
	}

	resp := response.DashboardMobileResponse{
		RondaSchedule: response.RondaScheduleModelToRondaScheduleResponse(schedules, nil),
		Announcecment: response.AnnouncementModelToAnnouncementResponse(announcement, driveService),
	}
	// get event limit 1
	return r.Ok(c, resp, "success", nil)
}
