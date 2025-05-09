package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type RondaConstributionHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type rondaConstributionHandlerImpl struct {
	services services.RondaConstributionService
}

func NewRondaConstributionHandler(services services.RondaConstributionService) RondaConstributionHandler {
	return &rondaConstributionHandlerImpl{services: services}
}

// Create RondaConstribution
// @Summary Create RondaConstribution
// @Descrpiton Create RondaConstribution
// @Tags Ronda Constribution
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaConstributionCreateRequest true "Create RondaConstribution"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-constribution [post]
func (h *rondaConstributionHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.RondaConstributionCreateRequest
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

// Update RondaConstribution
// @Summary Update RondaConstribution
// @Descrpiton Update RondaConstribution
// @Tags Ronda Constribution
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaConstributionUpdateRequset true "Update RondaConstribution"
// @Param id path string true "RondaConstribution id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-constribution/{id} [put]
func (h *rondaConstributionHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.RondaConstributionUpdateRequset)

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

// Get Pagination RondaConstribution
// @Summary Get Paginated RondaConstribution
// @Descrpiton Get Paginated RondaConstribution
// @Tags Ronda Constribution
// @Accept json
// @Produce json
// @Security Bearer
// @Param paginated query boolean false "Paginated"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order_by query string false "Order by"
// @Param order query string false "Order"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-constribution [get]
func (h *rondaConstributionHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.RondaConstributionResponses
	var err error

	if isPaginated {
		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate)
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
	} else {

		res, err := h.services.FindAll()
		if err != nil {
			return r.BadRequest(ctx, []string{err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find RondaConstribution By ID
// @Summary Find RondaConstribution By ID
// @Descrpiton Find RondaConstribution By ID
// @Tags Ronda Constribution
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaConstribution id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-constribution/{id} [get]
func (h *rondaConstributionHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete RondaConstribution By ID
// @Summary Delete RondaConstribution By ID
// @Descrpiton Delete RondaConstribution By ID
// @Tags Ronda Constribution
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaConstribution id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-constribution/{id} [delete]
func (h *rondaConstributionHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
