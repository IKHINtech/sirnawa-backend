package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplPaymentService interface {
	Create(data request.IplPaymentCreateRequest) (*response.IplPaymentResponse, error)
	Update(id string, data request.IplPaymentUpdateRequset) (*response.IplPaymentResponse, error)
	FindByID(id string) (*response.IplPaymentResponse, error)
	Delete(id string) error
	FindAll() (response.IplPaymentResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.IplPaymentResponses, error)
}

type iplPaymentServiceImpl struct {
	repository repository.IplPaymentRepository
	db         *gorm.DB
}

func NewIplPaymentServices(repo repository.IplPaymentRepository, db *gorm.DB) IplPaymentService {
	return &iplPaymentServiceImpl{repository: repo, db: db}
}

func (s *iplPaymentServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *iplPaymentServiceImpl) Create(data request.IplPaymentCreateRequest) (*response.IplPaymentResponse, error) {
	var result *models.IplPayment

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.IplPaymentCreateRequestToIplPaymentModel(data)

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

	res := response.IplPaymentModelToIplPaymentResponse(result)
	return res, nil
}

func (s *iplPaymentServiceImpl) Update(id string, data request.IplPaymentUpdateRequset) (*response.IplPaymentResponse, error) {
	var result *models.IplPayment

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.IplPaymentUpdateRequsetToIplPaymentModel(data)
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

	res := response.IplPaymentModelToIplPaymentResponse(result)
	return res, nil
}

func (s *iplPaymentServiceImpl) FindAll() (response.IplPaymentResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.IplPaymentListToResponse(result)
	return resp, nil
}

func (s *iplPaymentServiceImpl) FindByID(id string) (*response.IplPaymentResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.IplPaymentModelToIplPaymentResponse(result)
	return resp, err
}

func (s *iplPaymentServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.IplPaymentResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.IplPaymentListToResponse(data)
	return paginated, &resp, err
}

func (s *iplPaymentServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
