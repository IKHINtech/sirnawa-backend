package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplRateDetailService interface {
	Create(data request.IplRateDetailCreateRequest) (*response.IplRateDetailResponse, error)
	Update(id string, data request.IplRateDetailUpdateRequset) (*response.IplRateDetailResponse, error)
	FindByID(id string) (*response.IplRateDetailResponse, error)
	Delete(id string) error
	FindAll(iplRateID string) (response.IplRateDetailResponses, error)
	Paginated(pagination utils.Pagination, iplRateID string) (*utils.Pagination, *response.IplRateDetailResponses, error)
}

type iplRateDetailServiceImpl struct {
	repository repository.IplRateDetailRepository
	db         *gorm.DB
}

func NewIplRateDetailServices(repo repository.IplRateDetailRepository, db *gorm.DB) IplRateDetailService {
	return &iplRateDetailServiceImpl{repository: repo, db: db}
}

func (s *iplRateDetailServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *iplRateDetailServiceImpl) Create(data request.IplRateDetailCreateRequest) (*response.IplRateDetailResponse, error) {
	var result *models.IplRateDetail

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.IplRateDetailCreateRequestToIplRateDetailModel(data)

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

	res := response.IplRateDetailModelToIplRateDetailResponse(result)
	return res, nil
}

func (s *iplRateDetailServiceImpl) Update(id string, data request.IplRateDetailUpdateRequset) (*response.IplRateDetailResponse, error) {
	var result *models.IplRateDetail

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.IplRateDetailUpdateRequsetToIplRateDetailModel(data)
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

	res := response.IplRateDetailModelToIplRateDetailResponse(result)
	return res, nil
}

func (s *iplRateDetailServiceImpl) FindAll(iplRateID string) (response.IplRateDetailResponses, error) {
	result, err := s.repository.FindAll(iplRateID)
	if err != nil {
		return nil, err
	}

	resp := response.IplRateDetailListToResponse(result)
	return resp, nil
}

func (s *iplRateDetailServiceImpl) FindByID(id string) (*response.IplRateDetailResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.IplRateDetailModelToIplRateDetailResponse(result)
	return resp, err
}

func (s *iplRateDetailServiceImpl) Paginated(pagination utils.Pagination, iplRateID string) (*utils.Pagination, *response.IplRateDetailResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, iplRateID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.IplRateDetailListToResponse(data)
	return paginated, &resp, err
}

func (s *iplRateDetailServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
