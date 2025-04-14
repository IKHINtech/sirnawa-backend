package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BaseHandler interface {
	Create(ctx *fiber.Ctx)
	Update(ctx *fiber.Ctx)
	Paginated(ctx *fiber.Ctx)
	FindAll(ctx *fiber.Ctx)
	FindByID(ctx *fiber.Ctx)
	Delete(ctx *fiber.Ctx)
}

type baseHandlerImpl struct {
	services services.BaseService
}

func NewBaseHandler(services services.BaseService) BaseHandler {
	return &baseHandlerImpl{services: services}
}

func (h *baseHandlerImpl) Create(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	var req request.BaseCreateRequest
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

func (h *baseHandlerImpl) Update(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		r.BadRequest(ctx, []string{"id is required"})
		return
	}

	req := new(request.BaseUpdateRequset)

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

func (h *baseHandlerImpl) Paginated(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}

	paginate := utils.GetPaginationParams(ctx)

	meta, data, err := h.services.Paginated(paginate)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}
	r.Ok(ctx, data, "Successfully get data", meta)
}

func (h *baseHandlerImpl) FindAll(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	res, err := h.services.FindAll()
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
		return
	}
	r.Ok(ctx, res, "Successfully get data", nil)
}

func (h *baseHandlerImpl) FindByID(ctx *fiber.Ctx) {
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

func (h *baseHandlerImpl) Delete(ctx *fiber.Ctx) {
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
