package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type RondaGroupMemberHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type rondaGroupMemberHandlerImpl struct {
	services services.RondaGroupMemberService
}

func NewRondaGroupMemberHandler(services services.RondaGroupMemberService) RondaGroupMemberHandler {
	return &rondaGroupMemberHandlerImpl{services: services}
}

// Create RondaGroupMember
// @Summary Create RondaGroupMember
// @Descrpiton Create RondaGroupMember
// @Tags Ronda Group Member
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaGroupMemberCreateRequest true "Create RondaGroupMember"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group-member [post]
func (h *rondaGroupMemberHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.RondaGroupMemberCreateRequest
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

// Update RondaGroupMember
// @Summary Update RondaGroupMember
// @Descrpiton Update RondaGroupMember
// @Tags Ronda Group Member
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.RondaGroupMemberUpdateRequset true "Update RondaGroupMember"
// @Param id path string true "RondaGroupMember id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group-member/{id} [put]
func (h *rondaGroupMemberHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.RondaGroupMemberUpdateRequset)

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

// Get Pagination RondaGroupMember
// @Summary Get Paginated RondaGroupMember
// @Descrpiton Get Paginated RondaGroupMember
// @Tags Ronda Group Member
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
// @Router /ronda-group-member [get]
func (h *rondaGroupMemberHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.RondaGroupMemberResponses
	var err error

	if isPaginated {
		paginate := utils.GetPaginationParams(ctx)

		meta, data, err = h.services.Paginated(paginate)
		if err != nil {
			return r.BadRequest(ctx, []string{"error:" + err.Error()})
		}
	} else {

		res, err := h.services.FindAll()
		if err != nil {
			return r.BadRequest(ctx, []string{"error:" + err.Error()})
		}
		data = &res
	}
	return r.Ok(ctx, data, "Successfully get data", meta)
}

// Find RondaGroupMember By ID
// @Summary Find RondaGroupMember By ID
// @Descrpiton Find RondaGroupMember By ID
// @Tags Ronda Group Member
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaGroupMember id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group-member/{id} [get]
func (h *rondaGroupMemberHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete RondaGroupMember By ID
// @Summary Delete RondaGroupMember By ID
// @Descrpiton Delete RondaGroupMember By ID
// @Tags Ronda Group Member
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "RondaGroupMember id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /ronda-group-member/{id} [delete]
func (h *rondaGroupMemberHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
