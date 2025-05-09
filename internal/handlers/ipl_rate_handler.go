package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IplRateHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type iplRateHandlerImpl struct {
	services services.IplRateService
}

func NewIplRateHandler(services services.IplRateService) IplRateHandler {
	return &iplRateHandlerImpl{services: services}
}

// Create IplRate
// @Summary Create IplRate
// @Descrpiton Create IplRate
// @Tags Ipl Rate
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplRateCreateRequest true "Create IplRate"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate [post]
func (h *iplRateHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.IplRateCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid"})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Create(req)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Update IplRate
// @Summary Update IplRate
// @Descrpiton Update IplRate
// @Tags Ipl Rate
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplRateUpdateRequset true "Update IplRate"
// @Param id path string true "IplRate id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate/{id} [put]
func (h *iplRateHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.IplRateUpdateRequset)

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

// Get Pagination IplRate
// @Summary Get Paginated IplRate
// @Descrpiton Get Paginated IplRate
// @Tags Ipl Rate
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param rt_id query string false "RT ID"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate [get]
func (h *iplRateHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	rtID := ctx.Query("rt_id")
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.IplRateResponses
	var err error

	if isPaginated {
		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, rtID)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {

		res, err := h.services.FindAll(rtID)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find IplRate By ID
// @Summary Find IplRate By ID
// @Descrpiton Find IplRate By ID
// @Tags Ipl Rate
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplRate id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate/{id} [get]
func (h *iplRateHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete IplRate By ID
// @Summary Delete IplRate By ID
// @Descrpiton Delete IplRate By ID
// @Tags Ipl Rate
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplRate id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate/{id} [delete]
func (h *iplRateHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
