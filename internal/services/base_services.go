package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type BaseService interface {
	Create(data request.BaseCreateRequest) (*response.BaseResponse, error)
	Update(id string, data request.BaseUpdateRequset) (*response.BaseResponse, error)
	FindByID(id string) (*response.BaseResponse, error)
	Delete(id string) error
	FindAll() (response.BaseResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.BaseResponses, error)
}

type baseServiceImpl struct {
	repository repository.BaseRepository
	db         *gorm.DB
}

func NewBaseRepository(repo repository.BaseRepository, db *gorm.DB) BaseService {
	return &baseServiceImpl{repository: repo, db: db}
}

func (s *baseServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *baseServiceImpl) Create(data request.BaseCreateRequest) (*response.BaseResponse, error) {
	var result *models.BaseModel

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.BaseCreateRequestToBaseModel(data)

		created, err := s.repository.Create(tx, *payload)
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

	res := response.BaseModelToBaseResponse(result)
	return res, nil
}

func (s *baseServiceImpl) Update(id string, data request.BaseUpdateRequset) (*response.BaseResponse, error) {
	var result *models.BaseModel

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.BaseUpdateRequsetToBaseModel(data)
		payload.ID = existing.ID

		updated, err := s.repository.Update(tx, payload.ID, *payload)
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

	res := response.BaseModelToBaseResponse(result)
	return res, nil
}

func (s *baseServiceImpl) FindAll() (response.BaseResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.BaseListToResponse(result)
	return *resp, nil
}

func (s *baseServiceImpl) FindByID(id string) (*response.BaseResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.BaseModelToBaseResponse(result)
	return resp, err
}

func (s *baseServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.BaseResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.BaseListToResponse(data)
	return paginated, resp, err
}

func (s *baseServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
