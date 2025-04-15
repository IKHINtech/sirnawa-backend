package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type HouseService interface {
	Create(data request.HouseCreateRequest) (*response.HouseResponse, error)
	Update(id string, data request.HouseUpdateRequset) (*response.HouseResponse, error)
	FindByID(id string) (*response.HouseResponse, error)
	Delete(id string) error
	FindAll() (response.HouseResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.HouseResponses, error)
}

type houseServiceImpl struct {
	repository repository.HouseRepository
	db         *gorm.DB
}

func NewHouseServices(repo repository.HouseRepository, db *gorm.DB) HouseService {
	return &houseServiceImpl{repository: repo, db: db}
}

func (s *houseServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *houseServiceImpl) Create(data request.HouseCreateRequest) (*response.HouseResponse, error) {
	var result *models.House

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.HouseCreateRequestToHouseModel(data)

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

	res := response.HouseModelToHouseResponse(result)
	return res, nil
}

func (s *houseServiceImpl) Update(id string, data request.HouseUpdateRequset) (*response.HouseResponse, error) {
	var result *models.House

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.HouseUpdateRequsetToHouseModel(data)
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

	res := response.HouseModelToHouseResponse(result)
	return res, nil
}

func (s *houseServiceImpl) FindAll() (response.HouseResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.HouseListToResponse(result)
	return resp, nil
}

func (s *houseServiceImpl) FindByID(id string) (*response.HouseResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.HouseModelToHouseResponse(result)
	return resp, err
}

func (s *houseServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.HouseResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.HouseListToResponse(data)
	return paginated, &resp, err
}

func (s *houseServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
