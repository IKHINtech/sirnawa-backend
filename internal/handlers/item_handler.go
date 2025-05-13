package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ItemHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type itemHandlerImpl struct {
	services services.ItemService
}

func NewItemHandler(services services.ItemService) ItemHandler {
	return &itemHandlerImpl{services: services}
}

// Create Item
// @Summary Create Item
// @Descrpiton Create Item
// @Tags Item
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ItemCreateRequest true "Create Item"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /item [post]
func (h *itemHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.ItemCreateRequest
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

// Update Item
// @Summary Update Item
// @Descrpiton Update Item
// @Tags Item
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ItemUpdateRequset true "Update Item"
// @Param id path string true "Item id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /item/{id} [put]
func (h *itemHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.ItemUpdateRequset)

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

// Get Item
// @Summary Get Item
// @Descrpiton Get Item
// @Tags Item
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
// @Router /item [get]
func (h *itemHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rtID := ctx.Query("rt_id")
	isPaginated := ctx.QueryBool("paginated", true)

	if rtID == "" {
		r.Ok(ctx, response.ItemResponses{}, "RT ID harus di kirim", nil)
	}

	var meta *utils.Pagination
	var data *response.ItemResponses
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

// Find Item By ID
// @Summary Find Item By ID
// @Descrpiton Find Item By ID
// @Tags Item
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Item id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /item/{id} [get]
func (h *itemHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Item By ID
// @Summary Delete Item By ID
// @Descrpiton Delete Item By ID
// @Tags Item
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Item id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /item/{id} [delete]
func (h *itemHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
