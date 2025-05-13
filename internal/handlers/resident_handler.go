package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ResidentHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type residentHandlerImpl struct {
	services services.ResidentService
}

func NewResidentHandler(services services.ResidentService) ResidentHandler {
	return &residentHandlerImpl{services: services}
}

// Create Resident
// @Summary Create Resident
// @Descrpiton Create Resident
// @Tags Resident
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ResidentCreateRequest true "Create Resident"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident [post]
func (h *residentHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.ResidentCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid", err.Error()})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Create(req)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Update Resident
// @Summary Update Resident
// @Descrpiton Update Resident
// @Tags Resident
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ResidentUpdateRequset true "Update Resident"
// @Param id path string true "Resident id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident/{id} [put]
func (h *residentHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.ResidentUpdateRequset)

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

// Get Pagination Resident
// @Summary Get Paginated Resident
// @Descrpiton Get Paginated Resident
// @Tags Resident
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Param rt_id query string false "RT ID"
// @Param search query string false "Search Resident By Name"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident [get]
func (h *residentHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	search := ctx.Query("search", "")
	rt_id := ctx.Query("rt_id", "")
	isPaginated := ctx.QueryBool("paginated", true)

	if rt_id == "" {
		r.Ok(ctx, response.ResidentResponses{}, "RT ID harus di kirim", nil)
	}

	var meta *utils.Pagination
	var data *response.ResidentResponses
	var err error
	if isPaginated {
		paginate := utils.GetPaginationParams(ctx)
		paginate.SortBy = "residents.created_at"

		meta, data, err = h.services.Paginated(paginate, rt_id, search)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {
		res, err := h.services.FindAll(rt_id, search)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find Resident By ID
// @Summary Find Resident By ID
// @Descrpiton Find Resident By ID
// @Tags Resident
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Resident id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident/{id} [get]
func (h *residentHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Resident By ID
// @Summary Delete Resident By ID
// @Descrpiton Delete Resident By ID
// @Tags Resident
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Resident id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident/{id} [delete]
func (h *residentHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
