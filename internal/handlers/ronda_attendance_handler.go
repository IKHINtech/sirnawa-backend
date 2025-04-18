package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type RondaAttendanceHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type rondaAttendanceHandlerImpl struct {
	services services.RondaAttendanceService
}

func NewRondaAttendanceHandler(services services.RondaAttendanceService) RondaAttendanceHandler {
	return &rondaAttendanceHandlerImpl{services: services}
}

// Create RondaAttendance
// @Summary Create RondaAttendance
// @Descrpiton Create RondaAttendance
// @Tags Ronda Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaAttendanceCreateRequest true "Create RondaAttendance"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-attendance [post]
func (h *rondaAttendanceHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.RondaAttendanceCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid"})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Create(req)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Update RondaAttendance
// @Summary Update RondaAttendance
// @Descrpiton Update RondaAttendance
// @Tags Ronda Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaAttendanceUpdateRequset true "Update RondaAttendance"
// @Param id path string true "RondaAttendance id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-attendance/{id} [put]
func (h *rondaAttendanceHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.RondaAttendanceUpdateRequset)

	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid"})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Update(id, *req)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Get Pagination RondaAttendance
// @Summary Get Paginated RondaAttendance
// @Descrpiton Get Paginated RondaAttendance
// @Tags Ronda Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-attendance/paginated [get]
func (h *rondaAttendanceHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Get List RondaAttendance
// @Summary Get List RondaAttendance
// @Descrpiton Get List RondaAttendance
// @Tags Ronda Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-attendance [get]
func (h *rondaAttendanceHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Find RondaAttendance By ID
// @Summary Find RondaAttendance By ID
// @Descrpiton Find RondaAttendance By ID
// @Tags Ronda Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaAttendance id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-attendance/{id} [get]
func (h *rondaAttendanceHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	res, err := h.services.FindByID(id)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Delete RondaAttendance By ID
// @Summary Delete RondaAttendance By ID
// @Descrpiton Delete RondaAttendance By ID
// @Tags Ronda Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaAttendance id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-attendance/{id} [delete]
func (h *rondaAttendanceHandlerImpl) Delete(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	err := h.services.Delete(id)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, nil, "Successfully deleted", nil)
}
