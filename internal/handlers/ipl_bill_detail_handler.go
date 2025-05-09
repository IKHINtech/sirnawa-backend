package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IplBillDetailHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type iplBillDetailHandlerImpl struct {
	services services.IplBillDetailService
}

func NewIplBillDetailHandler(services services.IplBillDetailService) IplBillDetailHandler {
	return &iplBillDetailHandlerImpl{services: services}
}

// Create IplBillDetail
// @Summary Create Ipl Bill Detail
// @Descrpiton Create Ipl Bill Detail
// @Tags Ipl Bill Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplBillDetailCreateRequest true "Create IplBillDetail"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill-detail [post]
func (h *iplBillDetailHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.IplBillDetailCreateRequest
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

// Update Ipl Bill Detail
// @Summary Update Ipl Bill Detail
// @Descrpiton Update Ipl Bill Detail
// @Tags Ipl Bill Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplBillDetailUpdateRequset true "Update IplBillDetail"
// @Param id path string true "IplBillDetail id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill-detail/{id} [put]
func (h *iplBillDetailHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.IplBillDetailUpdateRequset)

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

// Get IplBillDetail
// @Summary Get Ipl Bill Detail
// @Descrpiton Get Ipl Bill Detail
// @Tags Ipl Bill Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param ipl_bill_id query string false "RT ID"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill-detail [get]
func (h *iplBillDetailHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	ipl_bill_id := ctx.Query("ipl_bill_id")
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.IplBillDetailResponses
	var err error

	if isPaginated {

		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, ipl_bill_id)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {
		res, err := h.services.FindAll(ipl_bill_id)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}

	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find Ipl Bill Detail By ID
// @Summary Find Ipl Bill Detail By ID
// @Descrpiton Find Ipl Bill Detail By ID
// @Tags Ipl Bill Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Ipl Bill Detail id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill-detail/{id} [get]
func (h *iplBillDetailHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Ipl Bill Detail By ID
// @Summary Delete Ipl Bill Detail By ID
// @Descrpiton Delete Ipl Bill Detail By ID
// @Tags Ipl Bill Detail
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Ipl Bill Detail id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill-detail/{id} [delete]
func (h *iplBillDetailHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
