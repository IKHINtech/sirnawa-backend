package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ShopHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type shopHandlerImpl struct {
	services services.ShopService
}

func NewShopHandler(services services.ShopService) ShopHandler {
	return &shopHandlerImpl{services: services}
}

// Create Shop
// @Summary Create Shop
// @Descrpiton Create Shop
// @Tags Shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ShopCreateRequest true "Create Shop"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop [post]
func (h *shopHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.ShopCreateRequest
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

// Update Shop
// @Summary Update Shop
// @Descrpiton Update Shop
// @Tags Shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ShopUpdateRequset true "Update Shop"
// @Param id path string true "Shop id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop/{id} [put]
func (h *shopHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.ShopUpdateRequset)

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

// Get Pagination Shop
// @Summary Get Paginated Shop
// @Descrpiton Get Paginated Shop
// @Tags Shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop/paginated [get]
func (h *shopHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Get List Shop
// @Summary Get List Shop
// @Descrpiton Get List Shop
// @Tags Shop
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop [get]
func (h *shopHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Find Shop By ID
// @Summary Find Shop By ID
// @Descrpiton Find Shop By ID
// @Tags Shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Shop id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop/{id} [get]
func (h *shopHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Shop By ID
// @Summary Delete Shop By ID
// @Descrpiton Delete Shop By ID
// @Tags Shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Shop id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop/{id} [delete]
func (h *shopHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
