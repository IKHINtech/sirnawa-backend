package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaActivityService interface {
	Create(data request.RondaActivityCreateRequest) (*response.RondaActivityResponse, error)
	Update(id string, data request.RondaActivityUpdateRequset) (*response.RondaActivityResponse, error)
	FindByID(id string) (*response.RondaActivityResponse, error)
	Delete(id string) error
	FindAll() (response.RondaActivityResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaActivityResponses, error)
}

type rondaActivityServiceImpl struct {
	repository repository.RondaActivityRepository
	db         *gorm.DB
}

func NewRondaActivityServices(repo repository.RondaActivityRepository, db *gorm.DB) RondaActivityService {
	return &rondaActivityServiceImpl{repository: repo, db: db}
}

func (s *rondaActivityServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rondaActivityServiceImpl) Create(data request.RondaActivityCreateRequest) (*response.RondaActivityResponse, error) {
	var result *models.RondaActivity

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RondaActivityCreateRequestToRondaActivityModel(data)

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

	res := response.RondaActivityModelToRondaActivityResponse(result)
	return res, nil
}

func (s *rondaActivityServiceImpl) Update(id string, data request.RondaActivityUpdateRequset) (*response.RondaActivityResponse, error) {
	var result *models.RondaActivity

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RondaActivityUpdateRequsetToRondaActivityModel(data)
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

	res := response.RondaActivityModelToRondaActivityResponse(result)
	return res, nil
}

func (s *rondaActivityServiceImpl) FindAll() (response.RondaActivityResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.RondaActivityListToResponse(result)
	return resp, nil
}

func (s *rondaActivityServiceImpl) FindByID(id string) (*response.RondaActivityResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RondaActivityModelToRondaActivityResponse(result)
	return resp, err
}

func (s *rondaActivityServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaActivityResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RondaActivityListToResponse(data)
	return paginated, &resp, err
}

func (s *rondaActivityServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
