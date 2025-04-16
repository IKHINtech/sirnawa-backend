package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaGroupMemberService interface {
	Create(data request.RondaGroupMemberCreateRequest) (*response.RondaGroupMemberResponse, error)
	Update(id string, data request.RondaGroupMemberUpdateRequset) (*response.RondaGroupMemberResponse, error)
	FindByID(id string) (*response.RondaGroupMemberResponse, error)
	Delete(id string) error
	FindAll() (response.RondaGroupMemberResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaGroupMemberResponses, error)
}

type rondaGroupMemberServiceImpl struct {
	repository repository.RondaGroupMemberRepository
	db         *gorm.DB
}

func NewRondaGroupMemberServices(repo repository.RondaGroupMemberRepository, db *gorm.DB) RondaGroupMemberService {
	return &rondaGroupMemberServiceImpl{repository: repo, db: db}
}

func (s *rondaGroupMemberServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rondaGroupMemberServiceImpl) Create(data request.RondaGroupMemberCreateRequest) (*response.RondaGroupMemberResponse, error) {
	var result *models.RondaGroupMember

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RondaGroupMemberCreateRequestToRondaGroupMemberModel(data)

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

	res := response.RondaGroupMemberModelToRondaGroupMemberResponse(result)
	return res, nil
}

func (s *rondaGroupMemberServiceImpl) Update(id string, data request.RondaGroupMemberUpdateRequset) (*response.RondaGroupMemberResponse, error) {
	var result *models.RondaGroupMember

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RondaGroupMemberUpdateRequsetToRondaGroupMemberModel(data)
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

	res := response.RondaGroupMemberModelToRondaGroupMemberResponse(result)
	return res, nil
}

func (s *rondaGroupMemberServiceImpl) FindAll() (response.RondaGroupMemberResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.RondaGroupMemberListToResponse(result)
	return resp, nil
}

func (s *rondaGroupMemberServiceImpl) FindByID(id string) (*response.RondaGroupMemberResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RondaGroupMemberModelToRondaGroupMemberResponse(result)
	return resp, err
}

func (s *rondaGroupMemberServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaGroupMemberResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RondaGroupMemberListToResponse(data)
	return paginated, &resp, err
}

func (s *rondaGroupMemberServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
