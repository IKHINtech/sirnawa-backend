package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplRateService interface {
	Create(data request.IplRateCreateRequest) (*response.IplRateResponse, error)
	Update(id string, data request.IplRateUpdateRequset) (*response.IplRateResponse, error)
	FindByID(id string) (*response.IplRateResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.IplRateResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.IplRateResponses, error)
}

type iplRateServiceImpl struct {
	repository repository.IplRateRepository
	db         *gorm.DB
}

func NewIplRateServices(repo repository.IplRateRepository, db *gorm.DB) IplRateService {
	return &iplRateServiceImpl{repository: repo, db: db}
}

func (s *iplRateServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *iplRateServiceImpl) Create(data request.IplRateCreateRequest) (*response.IplRateResponse, error) {
	var result *models.IplRate

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.IplRateCreateRequestToIplRateModel(data)

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

	res := response.IplRateModelToIplRateResponse(result)
	return res, nil
}

func (s *iplRateServiceImpl) Update(id string, data request.IplRateUpdateRequset) (*response.IplRateResponse, error) {
	var result *models.IplRate

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.IplRateUpdateRequsetToIplRateModel(data)
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

	res := response.IplRateModelToIplRateResponse(result)
	return res, nil
}

func (s *iplRateServiceImpl) FindAll(rtID string) (response.IplRateResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.IplRateListToResponse(result)
	return resp, nil
}

func (s *iplRateServiceImpl) FindByID(id string) (*response.IplRateResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.IplRateModelToIplRateResponse(result)
	return resp, err
}

func (s *iplRateServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.IplRateResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.IplRateListToResponse(data)
	return paginated, &resp, err
}

func (s *iplRateServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
