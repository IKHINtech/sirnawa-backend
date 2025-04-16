package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ShopService interface {
	Create(data request.ShopCreateRequest) (*response.ShopResponse, error)
	Update(id string, data request.ShopUpdateRequset) (*response.ShopResponse, error)
	FindByID(id string) (*response.ShopResponse, error)
	Delete(id string) error
	FindAll() (response.ShopResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ShopResponses, error)
}

type shopServiceImpl struct {
	repository repository.ShopRepository
	db         *gorm.DB
}

func NewShopServices(repo repository.ShopRepository, db *gorm.DB) ShopService {
	return &shopServiceImpl{repository: repo, db: db}
}

func (s *shopServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *shopServiceImpl) Create(data request.ShopCreateRequest) (*response.ShopResponse, error) {
	var result *models.Shop

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.ShopCreateRequestToShopModel(data)

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

	res := response.ShopModelToShopResponse(result)
	return res, nil
}

func (s *shopServiceImpl) Update(id string, data request.ShopUpdateRequset) (*response.ShopResponse, error) {
	var result *models.Shop

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.ShopUpdateRequsetToShopModel(data)
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

	res := response.ShopModelToShopResponse(result)
	return res, nil
}

func (s *shopServiceImpl) FindAll() (response.ShopResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.ShopListToResponse(result)
	return resp, nil
}

func (s *shopServiceImpl) FindByID(id string) (*response.ShopResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.ShopModelToShopResponse(result)
	return resp, err
}

func (s *shopServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ShopResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.ShopListToResponse(data)
	return paginated, &resp, err
}

func (s *shopServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
