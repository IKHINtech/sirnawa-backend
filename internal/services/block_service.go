package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type BlockService interface {
	Create(data request.BlockCreateRequest) (*response.BlockResponse, error)
	Update(id string, data request.BlockUpdateRequset) (*response.BlockResponse, error)
	FindByID(id string) (*response.BlockResponse, error)
	Delete(id string) error
	FindAll() (response.BlockResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.BlockResponses, error)
}

type blockServiceImpl struct {
	repository repository.BlockRepository
	db         *gorm.DB
}

func NewBlockRepository(repo repository.BlockRepository, db *gorm.DB) BlockService {
	return &blockServiceImpl{repository: repo, db: db}
}

func (s *blockServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *blockServiceImpl) Create(data request.BlockCreateRequest) (*response.BlockResponse, error) {
	var result *models.Block

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.BlockCreateRequestToBlockModel(data)

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

	res := response.BlockModelToBlockResponse(result)
	return res, nil
}

func (s *blockServiceImpl) Update(id string, data request.BlockUpdateRequset) (*response.BlockResponse, error) {
	var result *models.Block

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.BlockUpdateRequsetToBlockModel(data)
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

	res := response.BlockModelToBlockResponse(result)
	return res, nil
}

func (s *blockServiceImpl) FindAll() (response.BlockResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.BlockListToResponse(result)
	return resp, nil
}

func (s *blockServiceImpl) FindByID(id string) (*response.BlockResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.BlockModelToBlockResponse(result)
	return resp, err
}

func (s *blockServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.BlockResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.BlockListToResponse(data)
	return paginated, &resp, err
}

func (s *blockServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
