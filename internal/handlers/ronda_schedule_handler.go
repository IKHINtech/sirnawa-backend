package handlers

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type RondaScheduleHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type rondaScheduleHandlerImpl struct {
	services services.RondaScheduleService
}

func NewRondaScheduleHandler(services services.RondaScheduleService) RondaScheduleHandler {
	return &rondaScheduleHandlerImpl{services: services}
}

// Create RondaSchedule
// @Summary Create RondaSchedule
// @Descrpiton Create RondaSchedule
// @Tags Ronda Schedule
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaScheduleCreateRequest true "Create RondaSchedule"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-schedule [post]
func (h *rondaScheduleHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.RondaScheduleCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid, error: " + err.Error()})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Create(req)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Update RondaSchedule
// @Summary Update RondaSchedule
// @Descrpiton Update RondaSchedule
// @Tags Ronda Schedule
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaScheduleUpdateRequset true "Update RondaSchedule"
// @Param id path string true "RondaSchedule id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-schedule/{id} [put]
func (h *rondaScheduleHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.RondaScheduleUpdateRequset)

	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid"})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Update(id, *req)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Get Pagination RondaSchedule
// @Summary Get Paginated RondaSchedule
// @Descrpiton Get Paginated RondaSchedule
// @Tags Ronda Schedule
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param page query int false "Page number"
// @Param date query string false "Date (YYYY-MM-DD)"
// @Param rt_id query string false "RT ID"
// @Param group_id query string false "Group ID"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-schedule [get]
func (h *rondaScheduleHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rtID := ctx.Query("rt_id")
	groupID := ctx.Query("group_id")
	dateStr := ctx.Query("date")
	isPaginated := ctx.QueryBool("paginated", true)
	var date *time.Time

	if rtID == "" {
		r.Ok(ctx, response.RondaScheduleResponses{}, "RT ID harus di kirim", nil)
	}

	if dateStr != "" {
		layout := "2006-01-02"
		datecft, err := time.Parse(layout, dateStr)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		date = &datecft
	}

	var meta *utils.Pagination
	var data *response.RondaScheduleResponses
	var err error

	if isPaginated {
		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, rtID, groupID, date)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {

		res, err := h.services.FindAll(rtID, groupID, date)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find RondaSchedule By ID
// @Summary Find RondaSchedule By ID
// @Descrpiton Find RondaSchedule By ID
// @Tags Ronda Schedule
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaSchedule id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-schedule/{id} [get]
func (h *rondaScheduleHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	res, err := h.services.FindByID(id)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Delete RondaSchedule By ID
// @Summary Delete RondaSchedule By ID
// @Descrpiton Delete RondaSchedule By ID
// @Tags Ronda Schedule
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaSchedule id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-schedule/{id} [delete]
func (h *rondaScheduleHandlerImpl) Delete(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	err := h.services.Delete(id)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}
	return r.Ok(ctx, nil, "Successfully deleted", nil)
}
