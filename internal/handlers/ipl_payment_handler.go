package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IplPaymentHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type iplPaymentHandlerImpl struct {
	services services.IplPaymentService
}

func NewIplPaymentHandler(services services.IplPaymentService) IplPaymentHandler {
	return &iplPaymentHandlerImpl{services: services}
}

// Create IplPayment
// @Summary Create IplPayment
// @Descrpiton Create IplPayment
// @Tags IplPayment
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplPaymentCreateRequest true "Create IplPayment"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-payment [post]
func (h *iplPaymentHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.IplPaymentCreateRequest
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

// Update IplPayment
// @Summary Update IplPayment
// @Descrpiton Update IplPayment
// @Tags IplPayment
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplPaymentUpdateRequset true "Update IplPayment"
// @Param id path string true "IplPayment id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-payment/{id} [put]
func (h *iplPaymentHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.IplPaymentUpdateRequset)

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

// Get Pagination IplPayment
// @Summary Get Paginated IplPayment
// @Descrpiton Get Paginated IplPayment
// @Tags IplPayment
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-payment/paginated [get]
func (h *iplPaymentHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Get List IplPayment
// @Summary Get List IplPayment
// @Descrpiton Get List IplPayment
// @Tags IplPayment
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-payment [get]
func (h *iplPaymentHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, res, "Successfully get data", nil)
}

// Find IplPayment By ID
// @Summary Find IplPayment By ID
// @Descrpiton Find IplPayment By ID
// @Tags IplPayment
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplPayment id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-payment/{id} [get]
func (h *iplPaymentHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete IplPayment By ID
// @Summary Delete IplPayment By ID
// @Descrpiton Delete IplPayment By ID
// @Tags IplPayment
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplPayment id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-payment/{id} [delete]
func (h *iplPaymentHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
