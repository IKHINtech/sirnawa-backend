package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type NotificationHandler interface {
	CreateNotification(c *fiber.Ctx) error
	GetNotifications(c *fiber.Ctx) error
	MarkAsRead(c *fiber.Ctx) error
	MarkAllAsRead(c *fiber.Ctx) error
	DeleteNotification(c *fiber.Ctx) error
	GetUnreadCount(c *fiber.Ctx) error
}

type notificationHandlerImpl struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) NotificationHandler {
	return &notificationHandlerImpl{service: service}
}

// CreateNotification handler
// @Summary Create Notification
// @Description Create Notification
// @Tags Notification
// @Accept json
// @Security Bearer
// @Produce json
// @Param register body models.Notification true "Notification"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /notification [post]
func (h *notificationHandlerImpl) CreateNotification(c *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	// TODO: nanti diubah jadi request bukan model
	var input models.Notification
	if err := c.BodyParser(&input); err != nil {
		return r.BadRequest(c, []string{err.Error(), "Invalid request body"})
	}

	// Dapatkan userID dari auth middleware
	userID := c.Locals("user_id").(string)
	input.UserID = userID

	notification, err := h.service.CreateNotification(c.Context(), input)
	if err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Created(c, notification, "Successfully created")
}

// Get Pagination Notification
// @Summary Get Paginated Notification
// @Descrpiton Get Paginated Notification
// @Tags Notification
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param rt_id query string false "RT ID"
// @Param house_id query string false "House ID"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /notification [get]
func (h *notificationHandlerImpl) GetNotifications(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	rt_id := c.Query("rt_id", "")
	house_id := c.Query("house_id", "")
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(c)
	meta, notifications, err := h.service.GetUserNotifications(paginate, userID, rt_id, house_id)
	if err != nil {
		return r.BadRequest(c, []string{err.Error(), "Failed to get notifications"})
	}

	return r.Ok(c, notifications, "Successfully get data", meta)
}

// MarkAsRead godoc
// @Summary Mark notification as read
// @Description Mark a specific notification as read
// @Tags Notification
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Notification ID"
// @Success 200 {object} utils.ResponseData "Notification marked as read successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid notification ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 404 {object} utils.ErrorResponse "Notification not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to mark notification as read"
// @Router /notification/{id}/read [patch]
func (h *notificationHandlerImpl) MarkAsRead(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")
	r := &utils.ResponseHandler{}

	if err := h.service.MarkNotificationAsRead(c.Context(), id, userID); err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}
	return r.Ok(c, nil, "Successfully marked as read", nil)
}

// MarkAllAsRead godoc
// @Summary Mark all notifications as read
// @Description Mark all notifications as read for the authenticated user
// @Tags Notification
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData "All notifications marked as read successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Failed to mark notifications as read"
// @Router /notification/read-all [patch]
func (h *notificationHandlerImpl) MarkAllAsRead(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	r := &utils.ResponseHandler{}
	if err := h.service.MarkAllNotificationsAsRead(c.Context(), userID); err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Ok(c, nil, "Successfully marked all notifications as read", nil)
}

// DeleteNotification handler
// DeleteNotification godoc
// @Summary Delete a notification
// @Description Delete a specific notification
// @Tags Notification
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Notification ID"
// @Success 200 {object} utils.ResponseData "Notification deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid notification ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 404 {object} utils.ErrorResponse "Notification not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to delete notification"
// @Router /notification/{id} [delete]
func (h *notificationHandlerImpl) DeleteNotification(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	r := &utils.ResponseHandler{}
	if err := h.service.DeleteNotification(c.Context(), id, userID); err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Ok(c, nil, "Successfully deleted", nil)
}

// GetUnreadCount godoc
// @Summary Get unread notifications count
// @Description Get count of unread notifications for the authenticated user
// @Tags Notification
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData{data=int} "Unread count retrieved successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Failed to get unread count"
// @Router /notification/unread-count [get]
func (h *notificationHandlerImpl) GetUnreadCount(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	r := &utils.ResponseHandler{}
	count, err := h.service.GetUnreadCount(c.Context(), userID)
	if err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Ok(c, count, "Successfully get Count", nil)
}
