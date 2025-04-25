package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IplHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type iplHandlerImpl struct {
	services services.IplService
}

func NewIplHandler(services services.IplService) IplHandler {
	return &iplHandlerImpl{services: services}
}

// Create Ipl
// @Summary Create Ipl
// @Descrpiton Create Ipl
// @Tags Ipl
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplCreateRequest true "Create Ipl"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl [post]
func (h *iplHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.IplCreateRequest
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

// Update Ipl
// @Summary Update Ipl
// @Descrpiton Update Ipl
// @Tags Ipl
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplUpdateRequset true "Update Ipl"
// @Param id path string true "Ipl id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl/{id} [put]
func (h *iplHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.IplUpdateRequset)

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

// Get Ipl
// @Summary Get Ipl
// @Descrpiton Get Ipl
// @Tags Ipl
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param rt_id query string false "RT ID"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl [get]
func (h *iplHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rtID := ctx.Query("rt_id")
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.IplResponses
	var err error

	if isPaginated {

		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, rtID)
		if err != nil {
			return r.BadRequest(ctx, []string{"error:" + err.Error()})
		}
	} else {
		res, err := h.services.FindAll(rtID)
		if err != nil {
			return r.BadRequest(ctx, []string{"error:" + err.Error()})
		}
		data = &res
	}

	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find Ipl By ID
// @Summary Find Ipl By ID
// @Descrpiton Find Ipl By ID
// @Tags Ipl
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Ipl id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl/{id} [get]
func (h *iplHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Ipl By ID
// @Summary Delete Ipl By ID
// @Descrpiton Delete Ipl By ID
// @Tags Ipl
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Ipl id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl/{id} [delete]
func (h *iplHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
