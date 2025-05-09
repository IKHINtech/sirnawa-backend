package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ItemService interface {
	Create(data request.ItemCreateRequest) (*response.ItemResponse, error)
	Update(id string, data request.ItemUpdateRequset) (*response.ItemResponse, error)
	FindByID(id string) (*response.ItemResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.ItemResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.ItemResponses, error)
}

type itemServiceImpl struct {
	repository repository.ItemRepository
	db         *gorm.DB
}

func NewItemServices(repo repository.ItemRepository, db *gorm.DB) ItemService {
	return &itemServiceImpl{repository: repo, db: db}
}

func (s *itemServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *itemServiceImpl) Create(data request.ItemCreateRequest) (*response.ItemResponse, error) {
	var result *models.Item

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.ItemCreateRequestToItemModel(data)

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

	res := response.ItemModelToItemResponse(result)
	return res, nil
}

func (s *itemServiceImpl) Update(id string, data request.ItemUpdateRequset) (*response.ItemResponse, error) {
	var result *models.Item

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.ItemUpdateRequsetToItemModel(data)
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

	res := response.ItemModelToItemResponse(result)
	return res, nil
}

func (s *itemServiceImpl) FindAll(rtID string) (response.ItemResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.ItemListToResponse(result)
	return resp, nil
}

func (s *itemServiceImpl) FindByID(id string) (*response.ItemResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.ItemModelToItemResponse(result)
	return resp, err
}

func (s *itemServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.ItemResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.ItemListToResponse(data)
	return paginated, &resp, err
}

func (s *itemServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
