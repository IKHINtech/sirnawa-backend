package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IplRateDetailHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type iplRateDetailHandlerImpl struct {
	services services.IplRateDetailService
}

func NewIplRateDetailHandler(services services.IplRateDetailService) IplRateDetailHandler {
	return &iplRateDetailHandlerImpl{services: services}
}

// Create IplRateDetail
// @Summary Create IplRateDetail
// @Descrpiton Create IplRateDetail
// @Tags Ipl Rate Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplRateDetailCreateRequest true "Create IplRateDetail"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate-detail [post]
func (h *iplRateDetailHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.IplRateDetailCreateRequest
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

// Update IplRateDetail
// @Summary Update IplRateDetail
// @Descrpiton Update IplRateDetail
// @Tags Ipl Rate Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplRateDetailUpdateRequset true "Update IplRateDetail"
// @Param id path string true "IplRateDetail id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate-detail/{id} [put]
func (h *iplRateDetailHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.IplRateDetailUpdateRequset)

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

// Get Pagination IplRateDetail
// @Summary Get Paginated IplRateDetail
// @Descrpiton Get Paginated IplRateDetail
// @Tags Ipl Rate Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Param ipl_rate_id query string false "IPL Rate ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate-detail [get]
func (h *iplRateDetailHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	ipl_rate_id := ctx.Query("ipl_rate_id", "")
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.IplRateDetailResponses
	var err error

	if isPaginated {
		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, ipl_rate_id)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {

		res, err := h.services.FindAll(ipl_rate_id)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find IplRateDetail By ID
// @Summary Find IplRateDetail By ID
// @Descrpiton Find IplRateDetail By ID
// @Tags Ipl Rate Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplRateDetail id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate-detail/{id} [get]
func (h *iplRateDetailHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete IplRateDetail By ID
// @Summary Delete IplRateDetail By ID
// @Descrpiton Delete IplRateDetail By ID
// @Tags Ipl Rate Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplRateDetail id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-rate-detail/{id} [delete]
func (h *iplRateDetailHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
