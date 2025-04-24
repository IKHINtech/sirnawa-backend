package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AnnouncementHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Paginated(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type announcementHandlerImpl struct {
	services services.AnnouncementService
}

func NewAnnouncementHandler(services services.AnnouncementService) AnnouncementHandler {
	return &announcementHandlerImpl{services: services}
}

// Create Announcement
// @Summary Create Announcement
// @Descrpiton Create Announcement
// @Tags Announcement
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.AnnouncementCreateRequest true "Create Announcement"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /announcement [post]
func (h *announcementHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.AnnouncementCreateRequest
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

// Update Announcement
// @Summary Update Announcement
// @Descrpiton Update Announcement
// @Tags Announcement
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.AnnouncementUpdateRequset true "Update Announcement"
// @Param id path string true "Announcement id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /announcement/{id} [put]
func (h *announcementHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	req := new(request.AnnouncementUpdateRequset)

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

// Get Announcement
// @Summary Get Announcement
// @Descrpiton Get Announcement
// @Tags Announcement
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
// @Router /announcement [get]
func (h *announcementHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	isPaginated := ctx.QueryBool("paginated", true)

	var meta *utils.Pagination
	var data *response.AnnouncementResponses
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

// Find Announcement By ID
// @Summary Find Announcement By ID
// @Descrpiton Find Announcement By ID
// @Tags Announcement
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Announcement id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /announcement/{id} [get]
func (h *announcementHandlerImpl) FindByID(ctx *fiber.Ctx) error {
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

// Delete Announcement By ID
// @Summary Delete Announcement By ID
// @Descrpiton Delete Announcement By ID
// @Tags Announcement
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Announcement id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /announcement/{id} [delete]
func (h *announcementHandlerImpl) Delete(ctx *fiber.Ctx) error {
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
