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
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Create(req)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
	}

	r.Created(ctx, res, "Successfully created")
}

func (h *baseHandlerImpl) Update(ctx *fiber.Ctx) {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	req := new(request.BaseUpdateRequset)

	if err := ctx.BodyParser(&req); err != nil {
		r.BadRequest(ctx, []string{"Body is not valid"})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.Update(id, *req)
	if err != nil {
		r.BadRequest(ctx, []string{"error:" + err.Error()})
	}

	r.Created(ctx, res, "Successfully created")
}
func (h *baseHandlerImpl) Paginated(ctx *fiber.Ctx) {}
func (h *baseHandlerImpl) FindAll(ctx *fiber.Ctx)   {}
func (h *baseHandlerImpl) FindByID(ctx *fiber.Ctx)  {}
func (h *baseHandlerImpl) Delete(ctx *fiber.Ctx)    {}
