package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BlockHandler interface {
	Create(ctx *fiber.Ctx)
	Update(ctx *fiber.Ctx)
	Paginated(ctx *fiber.Ctx)
	FindAll(ctx *fiber.Ctx)
	FindByID(ctx *fiber.Ctx)
	Delete(ctx *fiber.Ctx)
}

type blockHandlerImpl struct {
	services services.BlockService
}

func NewBlockHandler(services services.BlockService) BlockHandler {
	return &blockHandlerImpl{services: services}
}

// Create Block
// @Summary Create Block
// @Descrpiton Create Block
// @Tags Block
// @Accept json
// @Produce json
// @Security
func (h *blockHandlerImpl) Create(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	var req request.BlockCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		r.BadRequest(ctx, []string{"Body is not valid"})
		return
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Create(req)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}

	r.Created(ctx, res, "Successfully created")
}

func (h *blockHandlerImpl) Update(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		r.BadRequest(ctx, []string{"id is required"})
		return
	}

	req := new(request.BlockUpdateRequset)

	if err := ctx.BodyParser(&req); err != nil {
		r.BadRequest(ctx, []string{"Body is not valid"})
		return
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Update(id, *req)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}

	r.Created(ctx, res, "Successfully created")
}

func (h *blockHandlerImpl) Paginated(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}
	r.Ok(ctx, data, "Successfully get data", meta)
}

func (h *blockHandlerImpl) FindAll(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}
	r.Ok(ctx, res, "Successfully get data", nil)
}

func (h *blockHandlerImpl) FindByID(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		r.BadRequest(ctx, []string{"id is required"})
		return
	}

	res, err := h.services.FindByID(id)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}
	r.Ok(ctx, res, "Successfully get data", nil)
}

func (h *blockHandlerImpl) Delete(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		r.BadRequest(ctx, []string{"id is required"})
		return
	}

	err := h.services.Delete(id)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}
	r.Ok(ctx, nil, "Successfully deleted", nil)
}
