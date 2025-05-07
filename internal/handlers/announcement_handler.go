package handlers

import (
	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
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
	services     services.AnnouncementService
	driveService utils.DriveService
}

func NewAnnouncementHandler(services services.AnnouncementService, driveService utils.DriveService) AnnouncementHandler {
	return &announcementHandlerImpl{services: services, driveService: driveService}
}

// Create Announcement
// @Summary Create Announcement with attachments
// @Description Create Announcement with file attachments
// @Tags Announcement
// @Accept multipart/form-data
// @Produce json
// @Security Bearer
// @Param title formData string true "Announcement title"
// @Param content formData string true "Announcement content"
// @Param created_by formData string true "Creator user ID"
// @Param rt_id formData string true "RT ID"
// @Param attachments formData []file false "Attachment files"
// @Success 201 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /announcement [post]
func (h *announcementHandlerImpl) Create(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	// Parse form data
	form, err := ctx.MultipartForm()
	if err != nil {
		return r.BadRequest(ctx, []string{"Failed to parse form data"})
	}

	// Bind form data to request struct
	req := request.AnnouncementCreateRequest{
		Title:     form.Value["title"][0],
		Content:   form.Value["content"][0],
		CreatedBy: form.Value["created_by"][0],
		RtID:      form.Value["rt_id"][0],
	}

	// Upload attachments to Google Drive
	var attachmentIDs []string
	if files := form.File["attachments"]; len(files) > 0 {
		for _, file := range files {
			fileID, err := h.driveService.UploadToDrive(file, config.AppConfig.DRIVE_FOLDER)
			if err != nil {
				return r.BadRequest(ctx, []string{"Failed to upload attachment: " + err.Error()})
			}
			attachmentIDs = append(attachmentIDs, fileID)
		}
	}
	req.Attachments = attachmentIDs

	// Create announcement
	res, err := h.services.Create(req)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}

	return r.Created(ctx, res, "Announcement created successfully")
}

// Update Announcement
// @Summary Update Announcement with attachments
// @Description Update Announcement with file attachments
// @Tags Announcement
// @Accept multipart/form-data
// @Produce json
// @Security Bearer
// @Param id path string true "Announcement ID"
// @Param title formData string false "Announcement title"
// @Param content formData string false "Announcement content"
// @Param attachments formData []file false "Attachment files"
// @Param delete_attachments formData []string false "Attachment IDs to delete"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ResponseData
// @Router /announcement/{id} [put]
func (h *announcementHandlerImpl) Update(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}
	id := ctx.Params("id")

	if id == "" {
		return r.BadRequest(ctx, []string{"ID is required"})
	}

	// Parse form data
	form, err := ctx.MultipartForm()
	if err != nil {
		return r.BadRequest(ctx, []string{"Failed to parse form data"})
	}

	// Bind form data to request struct
	req := request.AnnouncementCreateRequest{
		Title:   form.Value["title"][0],
		Content: form.Value["content"][0],
	}

	// Handle attachments
	var newAttachmentIDs []string
	if files := form.File["attachments"]; len(files) > 0 {
		for _, file := range files {
			fileID, err := h.driveService.UploadToDrive(file, config.AppConfig.DRIVE_FOLDER)
			if err != nil {
				return r.BadRequest(ctx, []string{"Failed to upload attachment: " + err.Error()})
			}
			newAttachmentIDs = append(newAttachmentIDs, fileID)
		}
	}
	req.Attachments = newAttachmentIDs

	// Validate request
	payload := request.AnnouncementUpdateRequset{
		ID:                        id,
		AnnouncementCreateRequest: req,
	}

	// Handle attachment deletions
	if deletes := form.Value["delete_attachments"]; len(deletes) > 0 {
		payload.DeleteAttachments = deletes
	}

	// Update announcement
	res, err := h.services.Update(id, payload)
	if err != nil {
		return r.BadRequest(ctx, []string{err.Error()})
	}

	return r.Ok(ctx, res, "Announcement updated successfully", nil)
}

// Get Announcement
// @Summary Get Announcement
// @Descrpiton Get Announcement
// @Tags Announcement
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
// @Router /announcement [get]
func (h *announcementHandlerImpl) Paginated(ctx *fiber.Ctx) error {
	r := &utils.ResponseHandler{}

	isPaginated := ctx.QueryBool("paginated", true)
	rtID := ctx.Query("rt_id")

	var meta *utils.Pagination
	var data *response.AnnouncementResponses
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
