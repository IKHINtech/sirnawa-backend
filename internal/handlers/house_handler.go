package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type HouseHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type houseHandlerImpl struct {
	services services.HouseService
}

func NewHouseHandler(services services.HouseService) HouseHandler {
	return &houseHandlerImpl{services: services}
}

// Create House
// @Summary Create House
// @Descrpiton Create House
// @Tags House
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.HouseCreateRequest true "Create House"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /house [post]
func (h *houseHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.HouseCreateRequest
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

// Update House
// @Summary Update House
// @Descrpiton Update House
// @Tags House
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.HouseUpdateRequset true "Update House"
// @Param id path string true "House id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /house/{id} [put]
func (h *houseHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.HouseUpdateRequset)

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

// Get Pagination House
// @Summary Get Paginated House
// @Descrpiton Get Paginated House
// @Tags House
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param rt_id query string false "RT ID"
// @Param block_id query string false "Block ID"
// @Param not_in_group_ronda query string false "Filter Not In Group Ronda"
// @Param status query string false "Status Rumah" Enums(aktif, tidak_aktif, kontrakan)
// @Param exclude_status query string false "Exclue Status Rumah" Enums(aktif, tidak_aktif, kontrakan)
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /house [get]
func (h *houseHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rt_id := ctx.Query("rt_id", "")
	block_id := ctx.Query("block_id", "")
	status := ctx.Query("status", "")
	isNotInGroupRonda := ctx.Query("not_in_group_ronda", "")
	isPaginated := ctx.QueryBool("paginated", true)
	excludeStatus := ctx.Query("exclude_status", "")

	if rt_id == "" {
		r.Ok(ctx, response.HouseResponses{}, "RT ID harus di kirim", nil)
	}
	var meta *utils.Pagination
	var data *[]response.HouseResponseDetail
	var err error
	if isPaginated {

		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate, rt_id, block_id, status, excludeStatus)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {
		res, err := h.services.FindAll(rt_id, block_id, status, isNotInGroupRonda, excludeStatus)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find House By ID
// @Summary Find House By ID
// @Descrpiton Find House By ID
// @Tags House
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "House id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /house/{id} [get]
func (h *houseHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete House By ID
// @Summary Delete House By ID
// @Descrpiton Delete House By ID
// @Tags House
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "House id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /house/{id} [delete]
func (h *houseHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
