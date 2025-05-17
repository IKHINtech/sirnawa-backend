package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type FCMTokenHandler struct {
	tokenService *services.FCMTokenService
}

func NewFCMTokenHandler(tokenService *services.FCMTokenService) *FCMTokenHandler {
	return &FCMTokenHandler{tokenService: tokenService}
}

// RegisterToken godoc
// @Summary Register FCM token
// @Description Register a new FCM token for push notifications
// @Tags FCM
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body request.RegisterTokenRequest true "Token registration data"
// @Success 201 {object} utils.ResponseHandler "Token registered successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or validation error"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Failed to register token"
// @Router /fcm/register [post]
func (h *FCMTokenHandler) RegisterToken(c *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	var req request.RegisterTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return r.BadRequest(c, []string{err.Error(), "Invalid body Request"})
	}

	middleware.ValidateRequest(req)

	// Dapatkan userID dari JWT atau session
	userID := c.Locals("user_id").(string)

	if err := h.tokenService.RegisterToken(
		userID,
		req.Token,
		req.DeviceID,
		req.DeviceType,
		req.AppVersion,
		req.OSVersion,
	); err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Created(c, nil, "Successfully Created")
}

// RemoveToken godoc
// @Summary Remove FCM token
// @Description Remove registered FCM token
// @Tags FCM
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body request.RemoveTokenRequest true "Token removal data"
// @Success 200 {object} map[string]bool "Token removed successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body or token not found"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Failed to remove token"
// @Router /fcm/remove [post]
func (h *FCMTokenHandler) RemoveToken(c *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	var req request.RemoveTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return r.BadRequest(c, []string{err.Error(), "Invalid body Request"})
	}

	if err := h.tokenService.RemoveToken(req.Token); err != nil {
		return r.BadRequest(c, []string{err.Error(), "Failed to remove token"})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
