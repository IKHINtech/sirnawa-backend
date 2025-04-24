package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type AnnouncementService interface {
	Create(data request.AnnouncementCreateRequest) (*response.AnnouncementResponse, error)
	Update(id string, data request.AnnouncementUpdateRequset) (*response.AnnouncementResponse, error)
	FindByID(id string) (*response.AnnouncementResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.AnnouncementResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.AnnouncementResponses, error)
}

type announcementServiceImpl struct {
	repository repository.AnnouncementRepository
	db         *gorm.DB
}

func NewAnnouncementServices(repo repository.AnnouncementRepository, db *gorm.DB) AnnouncementService {
	return &announcementServiceImpl{repository: repo, db: db}
}

func (s *announcementServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *announcementServiceImpl) Create(data request.AnnouncementCreateRequest) (*response.AnnouncementResponse, error) {
	var result *models.Announcement

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.AnnouncementCreateRequestToAnnouncementModel(data)

		created, err := s.repository.Create(tx, payload)
		if err != nil {
			return err
		}

		result = created
		return nil
	})
	// handling err
	if err != nil {
		return nil, err
	}

	res := response.AnnouncementModelToAnnouncementResponse(result)
	return res, nil
}

func (s *announcementServiceImpl) Update(id string, data request.AnnouncementUpdateRequset) (*response.AnnouncementResponse, error) {
	var result *models.Announcement

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.AnnouncementUpdateRequsetToAnnouncementModel(data)
		payload.ID = existing.ID

		updated, err := s.repository.Update(tx, payload.ID, payload)
		if err != nil {
			return err
		}

		result = updated
		return nil
	})
	// handle error
	if err != nil {
		return nil, err
	}

	res := response.AnnouncementModelToAnnouncementResponse(result)
	return res, nil
}

func (s *announcementServiceImpl) FindAll(rtID string) (response.AnnouncementResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.AnnouncementListToResponse(result)
	return resp, nil
}

func (s *announcementServiceImpl) FindByID(id string) (*response.AnnouncementResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.AnnouncementModelToAnnouncementResponse(result)
	return resp, err
}

func (s *announcementServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.AnnouncementResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.AnnouncementListToResponse(data)
	return paginated, &resp, err
}

func (s *announcementServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
