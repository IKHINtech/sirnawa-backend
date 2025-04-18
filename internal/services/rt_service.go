package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RtService interface {
	Create(data request.RtCreateRequest) (*response.RtResponse, error)
	Update(id string, data request.RtUpdateRequset) (*response.RtResponse, error)
	FindByID(id string) (*response.RtResponse, error)
	Delete(id string) error
	FindAll() (response.RtResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RtResponses, error)
}

type rtServiceImpl struct {
	repository repository.RtRepository
	db         *gorm.DB
}

func NewRtServices(repo repository.RtRepository, db *gorm.DB) RtService {
	return &rtServiceImpl{repository: repo, db: db}
}

func (s *rtServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rtServiceImpl) Create(data request.RtCreateRequest) (*response.RtResponse, error) {
	var result *models.Rt

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RtCreateRequestToRtModel(data)

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

	res := response.RtModelToRtResponse(result)
	return res, nil
}

func (s *rtServiceImpl) Update(id string, data request.RtUpdateRequset) (*response.RtResponse, error) {
	var result *models.Rt

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RtUpdateRequsetToRtModel(data)
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

	res := response.RtModelToRtResponse(result)
	return res, nil
}

func (s *rtServiceImpl) FindAll() (response.RtResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.RtListToResponse(result)
	return resp, nil
}

func (s *rtServiceImpl) FindByID(id string) (*response.RtResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RtModelToRtResponse(result)
	return resp, err
}

func (s *rtServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RtResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RtListToResponse(data)
	return paginated, &resp, err
}

func (s *rtServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
