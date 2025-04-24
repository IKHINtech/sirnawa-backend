package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaGroupService interface {
	Create(data request.RondaGroupCreateRequest) (*response.RondaGroupResponse, error)
	Update(id string, data request.RondaGroupUpdateRequset) (*response.RondaGroupResponse, error)
	FindByID(id string) (*response.RondaGroupResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.RondaGroupResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.RondaGroupResponses, error)
}

type rondaGroupServiceImpl struct {
	repository repository.RondaGroupRepository
	db         *gorm.DB
}

func NewRondaGroupServices(repo repository.RondaGroupRepository, db *gorm.DB) RondaGroupService {
	return &rondaGroupServiceImpl{repository: repo, db: db}
}

func (s *rondaGroupServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rondaGroupServiceImpl) Create(data request.RondaGroupCreateRequest) (*response.RondaGroupResponse, error) {
	var result *models.RondaGroup

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RondaGroupCreateRequestToRondaGroupModel(data)

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

	res := response.RondaGroupModelToRondaGroupResponse(result)
	return res, nil
}

func (s *rondaGroupServiceImpl) Update(id string, data request.RondaGroupUpdateRequset) (*response.RondaGroupResponse, error) {
	var result *models.RondaGroup

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RondaGroupUpdateRequsetToRondaGroupModel(data)
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

	res := response.RondaGroupModelToRondaGroupResponse(result)
	return res, nil
}

func (s *rondaGroupServiceImpl) FindAll(rtID string) (response.RondaGroupResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.RondaGroupListToResponse(result)
	return resp, nil
}

func (s *rondaGroupServiceImpl) FindByID(id string) (*response.RondaGroupResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RondaGroupModelToRondaGroupResponse(result)
	return resp, err
}

func (s *rondaGroupServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.RondaGroupResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RondaGroupListToResponse(data)
	return paginated, &resp, err
}

func (s *rondaGroupServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
