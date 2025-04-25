package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BlockHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type blockHandlerImpl struct {
	services services.BlockService
}

func NewBlockHandler(services services.BlockService) BlockHandler {
	return &blockHandlerImpl{services: services}
}

// Create Block
// @Summary Create Block
// @Descrpiton Create Block
// @Tags Block
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.BlockCreateRequest true "Create Block"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /block [post]
func (h *blockHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.BlockCreateRequest
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

// Update Block
// @Summary Update Block
// @Descrpiton Update Block
// @Tags Block
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.BlockUpdateRequset true "Update Block"
// @Param id path string true "Block id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /block/{id} [put]
func (h *blockHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.BlockUpdateRequset)

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

// Get Block
// @Summary Get Block
// @Descrpiton Get Block
// @Tags Block
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
// @Router /block [get]
func (h *blockHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rtID := ctx.Query("rt_id")
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.BlockResponses
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

// Find Block By ID
// @Summary Find Block By ID
// @Descrpiton Find Block By ID
// @Tags Block
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Block id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /block/{id} [get]
func (h *blockHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Block By ID
// @Summary Delete Block By ID
// @Descrpiton Delete Block By ID
// @Tags Block
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Block id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /block/{id} [delete]
func (h *blockHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
