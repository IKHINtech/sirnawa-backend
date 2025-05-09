package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplBillDetailService interface {
	Create(data request.IplBillDetailCreateRequest) (*response.IplBillDetailResponse, error)
	Update(id string, data request.IplBillDetailUpdateRequset) (*response.IplBillDetailResponse, error)
	FindByID(id string) (*response.IplBillDetailResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.IplBillDetailResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.IplBillDetailResponses, error)
}

type iplBillDetailServiceImpl struct {
	repository repository.IplBillDetailRepository
	db         *gorm.DB
}

func NewIplBillDetailServices(repo repository.IplBillDetailRepository, db *gorm.DB) IplBillDetailService {
	return &iplBillDetailServiceImpl{repository: repo, db: db}
}

func (s *iplBillDetailServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *iplBillDetailServiceImpl) Create(data request.IplBillDetailCreateRequest) (*response.IplBillDetailResponse, error) {
	var result *models.IplBillDetail

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.IplBillDetailCreateRequestToIplBillDetailModel(data)

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

	res := response.IplBillDetailModelToIplBillDetailResponse(result)
	return res, nil
}

func (s *iplBillDetailServiceImpl) Update(id string, data request.IplBillDetailUpdateRequset) (*response.IplBillDetailResponse, error) {
	var result *models.IplBillDetail

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.IplBillDetailUpdateRequsetToIplBillDetailModel(data)
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

	res := response.IplBillDetailModelToIplBillDetailResponse(result)
	return res, nil
}

func (s *iplBillDetailServiceImpl) FindAll(rtID string) (response.IplBillDetailResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.IplBillDetailListToResponse(result)
	return resp, nil
}

func (s *iplBillDetailServiceImpl) FindByID(id string) (*response.IplBillDetailResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.IplBillDetailModelToIplBillDetailResponse(result)
	return resp, err
}

func (s *iplBillDetailServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.IplBillDetailResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.IplBillDetailListToResponse(data)
	return paginated, &resp, err
}

func (s *iplBillDetailServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
