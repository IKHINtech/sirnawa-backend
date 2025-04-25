package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type PostHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type postHandlerImpl struct {
	services services.PostService
}

func NewPostHandler(services services.PostService) PostHandler {
	return &postHandlerImpl{services: services}
}

// Create Post
// @Summary Create Post
// @Descrpiton Create Post
// @Tags Post
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.PostCreateRequest true "Create Post"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /post [post]
func (h *postHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.PostCreateRequest
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

// Update Post
// @Summary Update Post
// @Descrpiton Update Post
// @Tags Post
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.PostUpdateRequset true "Update Post"
// @Param id path string true "Post id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /post/{id} [put]
func (h *postHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.PostUpdateRequset)

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

// Get Pagination Post
// @Summary Get Paginated Post
// @Descrpiton Get Paginated Post
// @Tags Post
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
// @Router /post [get]
func (h *postHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	isPaginated := ctx.QueryBool("paginated", true)
	rtID := ctx.Query("rt_id")
	var meta *utils.Pagination
	var data *response.PostResponses
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

// Find Post By ID
// @Summary Find Post By ID
// @Descrpiton Find Post By ID
// @Tags Post
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Post id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /post/{id} [get]
func (h *postHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Post By ID
// @Summary Delete Post By ID
// @Descrpiton Delete Post By ID
// @Tags Post
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Post id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /post/{id} [delete]
func (h *postHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
