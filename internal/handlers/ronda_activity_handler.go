package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type RondaActivityHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type rondaActivityHandlerImpl struct {
	services services.RondaActivityService
}

func NewRondaActivityHandler(services services.RondaActivityService) RondaActivityHandler {
	return &rondaActivityHandlerImpl{services: services}
}

// Create RondaActivity
// @Summary Create RondaActivity
// @Descrpiton Create RondaActivity
// @Tags Ronda Activity
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaActivityCreateRequest true "Create RondaActivity"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-activity [post]
func (h *rondaActivityHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.RondaActivityCreateRequest
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

// Update RondaActivity
// @Summary Update RondaActivity
// @Descrpiton Update RondaActivity
// @Tags Ronda Activity
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaActivityUpdateRequset true "Update RondaActivity"
// @Param id path string true "RondaActivity id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-activity/{id} [put]
func (h *rondaActivityHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.RondaActivityUpdateRequset)

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

// Get Pagination RondaActivity
// @Summary Get Paginated RondaActivity
// @Descrpiton Get Paginated RondaActivity
// @Tags Ronda Activity
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-activity/paginated [get]
func (h *rondaActivityHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Get List RondaActivity
// @Summary Get List RondaActivity
// @Descrpiton Get List RondaActivity
// @Tags Ronda Activity
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-activity [get]
func (h *rondaActivityHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Find RondaActivity By ID
// @Summary Find RondaActivity By ID
// @Descrpiton Find RondaActivity By ID
// @Tags Ronda Activity
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaActivity id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-activity/{id} [get]
func (h *rondaActivityHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete RondaActivity By ID
// @Summary Delete RondaActivity By ID
// @Descrpiton Delete RondaActivity By ID
// @Tags Ronda Activity
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaActivity id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-activity/{id} [delete]
func (h *rondaActivityHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
