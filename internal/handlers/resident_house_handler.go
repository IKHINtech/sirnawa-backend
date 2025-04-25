package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ResidentHouseHandler interface {
	AssignResidentToHouse(ctx *fiber.Ctx) error
	ChangeToPrimary(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type residentHouseHandlerImpl struct {
	services services.ResidentHouseService
}

func NewResidentHouseHandler(services services.ResidentHouseService) ResidentHouseHandler {
	return &residentHouseHandlerImpl{services: services}
}

// Create ResidentHouse
// @Summary Create ResidentHouse
// @Descrpiton Create ResidentHouse
// @Tags Resident House
// @Accept json
// @Produce json
// @Security Bearer
// @Param data body request.ResidentHouseCreateRequest true "Create ResidentHouse"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident-house [post]
func (h *residentHouseHandlerImpl) AssignResidentToHouse(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	var req request.ResidentHouseCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return r.BadRequest(ctx, []string{"Body is not valid"})
	}

	middleware.ValidateRequest(req)

	res, err := h.services.AssignResidentToHouse(req)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}

	return r.Created(ctx, res, "Successfully created")
}

// Delete ResidentHouse By ID
// @Summary Delete ResidentHouse By ID
// @Descrpiton Delete ResidentHouse By ID
// @Tags Resident House
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ResidentHouse id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident-house/{id} [delete]
func (h *residentHouseHandlerImpl) Delete(ctx *fiber.Ctx) error {
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

// Change To Primary By ID
// @Summary Change To Primary By ID
// @Descrpiton Change To Primary By ID
// @Tags Resident House
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ResidentHouse id"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /resident-house/{id} [get]
func (h *residentHouseHandlerImpl) ChangeToPrimary(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")
	if id == "" {
		return r.BadRequest(ctx, []string{"id is required"})
	}

	err := h.services.ChangeToPrimary(id)
	if err != nil {
		return r.BadRequest(ctx, []string{"error:" + err.Error()})
	}
	return r.Ok(ctx, nil, "Successfully deleted", nil)
}
