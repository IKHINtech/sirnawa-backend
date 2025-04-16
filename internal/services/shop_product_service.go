package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ShopProductService interface {
	Create(data request.ShopProductCreateRequest) (*response.ShopProductResponse, error)
	Update(id string, data request.ShopProductUpdateRequset) (*response.ShopProductResponse, error)
	FindByID(id string) (*response.ShopProductResponse, error)
	Delete(id string) error
	FindAll() (response.ShopProductResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ShopProductResponses, error)
}

type shopProductServiceImpl struct {
	repository repository.ShopProductRepository
	db         *gorm.DB
}

func NewShopProductServices(repo repository.ShopProductRepository, db *gorm.DB) ShopProductService {
	return &shopProductServiceImpl{repository: repo, db: db}
}

func (s *shopProductServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *shopProductServiceImpl) Create(data request.ShopProductCreateRequest) (*response.ShopProductResponse, error) {
	var result *models.ShopProduct

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.ShopProductCreateRequestToShopProductModel(data)

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

	res := response.ShopProductModelToShopProductResponse(result)
	return res, nil
}

func (s *shopProductServiceImpl) Update(id string, data request.ShopProductUpdateRequset) (*response.ShopProductResponse, error) {
	var result *models.ShopProduct

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.ShopProductUpdateRequsetToShopProductModel(data)
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

	res := response.ShopProductModelToShopProductResponse(result)
	return res, nil
}

func (s *shopProductServiceImpl) FindAll() (response.ShopProductResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.ShopProductListToResponse(result)
	return resp, nil
}

func (s *shopProductServiceImpl) FindByID(id string) (*response.ShopProductResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.ShopProductModelToShopProductResponse(result)
	return resp, err
}

func (s *shopProductServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ShopProductResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.ShopProductListToResponse(data)
	return paginated, &resp, err
}

func (s *shopProductServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
