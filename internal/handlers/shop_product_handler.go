package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ShopProductHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type shopProductHandlerImpl struct {
	services services.ShopProductService
}

func NewShopProductHandler(services services.ShopProductService) ShopProductHandler {
	return &shopProductHandlerImpl{services: services}
}

// Create ShopProduct
// @Summary Create Shop Product
// @Descrpiton Create Shop Product
// @Tags Shop Product
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ShopProductCreateRequest true "Create ShopProduct"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop-product [post]
func (h *shopProductHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.ShopProductCreateRequest
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

// Update ShopProduct
// @Summary Update ShopProduct
// @Descrpiton Update ShopProduct
// @Tags Shop Product
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ShopProductUpdateRequset true "Update ShopProduct"
// @Param id path string true "ShopProduct id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop-product/{id} [put]
func (h *shopProductHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.ShopProductUpdateRequset)

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

// Get Pagination ShopProduct
// @Summary Get Paginated ShopProduct
// @Descrpiton Get Paginated ShopProduct
// @Tags Shop Product
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop-product/paginated [get]
func (h *shopProductHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Get List ShopProduct
// @Summary Get List ShopProduct
// @Descrpiton Get List ShopProduct
// @Tags Shop Product
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop-product [get]
func (h *shopProductHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Find ShopProduct By ID
// @Summary Find ShopProduct By ID
// @Descrpiton Find ShopProduct By ID
// @Tags Shop Product
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ShopProduct id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop-product/{id} [get]
func (h *shopProductHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete ShopProduct By ID
// @Summary Delete ShopProduct By ID
// @Descrpiton Delete ShopProduct By ID
// @Tags Shop Product
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ShopProduct id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /shop-product/{id} [delete]
func (h *shopProductHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
