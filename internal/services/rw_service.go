package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RwService interface {
	Create(data request.RwCreateRequest) (*response.RwResponse, error)
	Update(id string, data request.RwUpdateRequset) (*response.RwResponse, error)
	FindByID(id string) (*response.RwResponse, error)
	Delete(id string) error
	FindAll() (response.RwResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RwResponses, error)
}

type rwServiceImpl struct {
	repository repository.RwRepository
	db         *gorm.DB
}

func NewRwServices(repo repository.RwRepository, db *gorm.DB) RwService {
	return &rwServiceImpl{repository: repo, db: db}
}

func (s *rwServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rwServiceImpl) Create(data request.RwCreateRequest) (*response.RwResponse, error) {
	var result *models.Rw

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RwCreateRequestToRwModel(data)

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

	res := response.RwModelToRwResponse(result)
	return res, nil
}

func (s *rwServiceImpl) Update(id string, data request.RwUpdateRequset) (*response.RwResponse, error) {
	var result *models.Rw

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RwUpdateRequsetToRwModel(data)
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

	res := response.RwModelToRwResponse(result)
	return res, nil
}

func (s *rwServiceImpl) FindAll() (response.RwResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.RwListToResponse(result)
	return resp, nil
}

func (s *rwServiceImpl) FindByID(id string) (*response.RwResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RwModelToRwResponse(result)
	return resp, err
}

func (s *rwServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RwResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RwListToResponse(data)
	return paginated, &resp, err
}

func (s *rwServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
