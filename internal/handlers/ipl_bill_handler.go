package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IplBillHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type iplBillHandlerImpl struct {
	services services.IplBillService
}

func NewIplBillHandler(services services.IplBillService) IplBillHandler {
	return &iplBillHandlerImpl{services: services}
}

// Create IplBill
// @Summary Create IplBill
// @Descrpiton Create IplBill
// @Tags Ipl Bill
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplBillCreateRequest true "Create IplBill"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill [post]
func (h *iplBillHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.IplBillCreateRequest
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

// Update IplBill
// @Summary Update IplBill
// @Descrpiton Update IplBill
// @Tags Ipl Bill
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.IplBillUpdateRequset true "Update IplBill"
// @Param id path string true "IplBill id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill/{id} [put]
func (h *iplBillHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.IplBillUpdateRequset)

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

// Get IplBill
// @Summary Get IplBill
// @Descrpiton Get IplBill
// @Tags Ipl Bill
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param rt_id query string false "RT ID"
// @Param house_id query string false "House ID"
// @Param status query string false "Status"
// @Param month query int false "Bulan"
// @Param year query int false "Tahun"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill [get]
func (h *iplBillHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rtID := ctx.Query("rt_id")
	houseID := ctx.Query("house_id")
	month := ctx.QueryInt("month")
	year := ctx.QueryInt("year")
	status := ctx.Query("status")
	isPaginated := ctx.QueryBool("paginated", true)

	if rtID == "" {
		r.Ok(ctx, response.IplBillResponses{}, "RT ID harus di kirim", nil)
	}
	var meta *utils.Pagination
	var data *response.IplBillResponses
	var err error

	if isPaginated {

		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, rtID, houseID, status, month, year)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {
		res, err := h.services.FindAll(rtID, houseID, status, month, year)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}

	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find IplBill By ID
// @Summary Find IplBill By ID
// @Descrpiton Find IplBill By ID
// @Tags Ipl Bill
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplBill id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill/{id} [get]
func (h *iplBillHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete IplBill By ID
// @Summary Delete IplBill By ID
// @Descrpiton Delete IplBill By ID
// @Tags Ipl Bill
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "IplBill id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ipl-bill/{id} [delete]
func (h *iplBillHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
