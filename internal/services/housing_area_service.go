package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type HousingAreaService interface {
	Create(data request.HousingAreaCreateRequest) (*response.HousingAreaResponse, error)
	Update(id string, data request.HousingAreaUpdateRequset) (*response.HousingAreaResponse, error)
	FindByID(id string) (*response.HousingAreaResponse, error)
	Delete(id string) error
	FindAll() (response.HousingAreaResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.HousingAreaResponses, error)
}

type housingAreaServiceImpl struct {
	repository repository.HousingAreaRepository
	db         *gorm.DB
}

func NewHousingAreaServices(repo repository.HousingAreaRepository, db *gorm.DB) HousingAreaService {
	return &housingAreaServiceImpl{repository: repo, db: db}
}

func (s *housingAreaServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *housingAreaServiceImpl) Create(data request.HousingAreaCreateRequest) (*response.HousingAreaResponse, error) {
	var result *models.HousingArea

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.HousingAreaCreateRequestToHousingAreaModel(data)

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

	res := response.HousingAreaModelToHousingAreaResponse(result)
	return res, nil
}

func (s *housingAreaServiceImpl) Update(id string, data request.HousingAreaUpdateRequset) (*response.HousingAreaResponse, error) {
	var result *models.HousingArea

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.HousingAreaUpdateRequsetToHousingAreaModel(data)
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

	res := response.HousingAreaModelToHousingAreaResponse(result)
	return res, nil
}

func (s *housingAreaServiceImpl) FindAll() (response.HousingAreaResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.HousingAreaListToResponse(result)
	return resp, nil
}

func (s *housingAreaServiceImpl) FindByID(id string) (*response.HousingAreaResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.HousingAreaModelToHousingAreaResponse(result)
	return resp, err
}

func (s *housingAreaServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.HousingAreaResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.HousingAreaListToResponse(data)
	return paginated, &resp, err
}

func (s *housingAreaServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
