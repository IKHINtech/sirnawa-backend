package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaConstributionService interface {
	Create(data request.RondaConstributionCreateRequest) (*response.RondaConstributionResponse, error)
	Update(id string, data request.RondaConstributionUpdateRequset) (*response.RondaConstributionResponse, error)
	FindByID(id string) (*response.RondaConstributionResponse, error)
	Delete(id string) error
	FindAll() (response.RondaConstributionResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaConstributionResponses, error)
}

type rondaConstributtionServiceImpl struct {
	repository repository.RondaConstributionRepository
	db         *gorm.DB
}

func NewRondaConstributionServices(repo repository.RondaConstributionRepository, db *gorm.DB) RondaConstributionService {
	return &rondaConstributtionServiceImpl{repository: repo, db: db}
}

func (s *rondaConstributtionServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rondaConstributtionServiceImpl) Create(data request.RondaConstributionCreateRequest) (*response.RondaConstributionResponse, error) {
	var result *models.RondaConstribution

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RondaConstributionCreateRequestToRondaConstributionModel(data)

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

	res := response.RondaConstributionModelToRondaConstributionResponse(result)
	return res, nil
}

func (s *rondaConstributtionServiceImpl) Update(id string, data request.RondaConstributionUpdateRequset) (*response.RondaConstributionResponse, error) {
	var result *models.RondaConstribution

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RondaConstributionUpdateRequsetToRondaConstributionModel(data)
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

	res := response.RondaConstributionModelToRondaConstributionResponse(result)
	return res, nil
}

func (s *rondaConstributtionServiceImpl) FindAll() (response.RondaConstributionResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.RondaConstributionListToResponse(result)
	return resp, nil
}

func (s *rondaConstributtionServiceImpl) FindByID(id string) (*response.RondaConstributionResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RondaConstributionModelToRondaConstributionResponse(result)
	return resp, err
}

func (s *rondaConstributtionServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaConstributionResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RondaConstributionListToResponse(data)
	return paginated, &resp, err
}

func (s *rondaConstributtionServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
