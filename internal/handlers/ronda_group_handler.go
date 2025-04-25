package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type RondaGroupHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type rondaGroupHandlerImpl struct {
	services services.RondaGroupService
}

func NewRondaGroupHandler(services services.RondaGroupService) RondaGroupHandler {
	return &rondaGroupHandlerImpl{services: services}
}

// Create RondaGroup
// @Summary Create RondaGroup
// @Descrpiton Create RondaGroup
// @Tags Ronda Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaGroupCreateRequest true "Create RondaGroup"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group [post]
func (h *rondaGroupHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.RondaGroupCreateRequest
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

// Update RondaGroup
// @Summary Update RondaGroup
// @Descrpiton Update RondaGroup
// @Tags Ronda Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaGroupUpdateRequset true "Update RondaGroup"
// @Param id path string true "RondaGroup id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group/{id} [put]
func (h *rondaGroupHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.RondaGroupUpdateRequset)

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

// Get Pagination RondaGroup
// @Summary Get Paginated RondaGroup
// @Descrpiton Get Paginated RondaGroup
// @Tags Ronda Group
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
// @Router /ronda-group [get]
func (h *rondaGroupHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	rtID := ctx.Query("rt_id")
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.RondaGroupResponses
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

// Find RondaGroup By ID
// @Summary Find RondaGroup By ID
// @Descrpiton Find RondaGroup By ID
// @Tags Ronda Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaGroup id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group/{id} [get]
func (h *rondaGroupHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete RondaGroup By ID
// @Summary Delete RondaGroup By ID
// @Descrpiton Delete RondaGroup By ID
// @Tags Ronda Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaGroup id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group/{id} [delete]
func (h *rondaGroupHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
