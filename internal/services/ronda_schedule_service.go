package services

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaScheduleService interface {
	Create(data request.RondaScheduleCreateRequest) (*response.RondaScheduleResponse, error)
	Update(id string, data request.RondaScheduleUpdateRequset) (*response.RondaScheduleResponse, error)
	FindByID(id string) (*response.RondaScheduleResponse, error)
	Delete(id string) error
	FindAll(rtID, groupID string, date *time.Time) (response.RondaScheduleResponses, error)
	Paginated(pagination utils.Pagination, rtID, groupID string, date *time.Time) (*utils.Pagination, *response.RondaScheduleResponses, error)
}

type rondaGroupScheduleServiceImpl struct {
	repository            repository.RondaScheduleRepository
	groupMemberRepository repository.RondaGroupMemberRepository
	db                    *gorm.DB
}

func NewRondaScheduleServices(repo repository.RondaScheduleRepository,
	groupMemberRepository repository.RondaGroupMemberRepository,
	db *gorm.DB,
) RondaScheduleService {
	return &rondaGroupScheduleServiceImpl{repository: repo, groupMemberRepository: groupMemberRepository, db: db}
}

func (s *rondaGroupScheduleServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rondaGroupScheduleServiceImpl) Create(data request.RondaScheduleCreateRequest) (*response.RondaScheduleResponse, error) {
	var result *models.RondaSchedule

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RondaScheduleCreateRequestToRondaScheduleModel(data)

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

	res := response.RondaScheduleModelToRondaScheduleResponse(result, nil)
	return res, nil
}

func (s *rondaGroupScheduleServiceImpl) Update(id string, data request.RondaScheduleUpdateRequset) (*response.RondaScheduleResponse, error) {
	var result *models.RondaSchedule

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RondaScheduleUpdateRequsetToRondaScheduleModel(data)
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

	res := response.RondaScheduleModelToRondaScheduleResponse(result, nil)
	return res, nil
}

func (s *rondaGroupScheduleServiceImpl) FindAll(rtID, groupID string, date *time.Time) (response.RondaScheduleResponses, error) {
	data, err := s.repository.FindAll(rtID, groupID, date)
	if err != nil {
		return nil, err
	}

	resp := make(response.RondaScheduleResponses, len(data))
	for i, item := range data {
		totalMember, err := s.groupMemberRepository.GetTotalMember(item.GroupID)
		if err != nil {
			return nil, err
		}
		resp[i] = *response.RondaScheduleModelToRondaScheduleResponse(&item, totalMember)
	}

	return resp, nil
}

func (s *rondaGroupScheduleServiceImpl) FindByID(id string) (*response.RondaScheduleResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RondaScheduleModelToRondaScheduleResponse(result, nil)
	return resp, err
}

func (s *rondaGroupScheduleServiceImpl) Paginated(pagination utils.Pagination, rtID, groupID string, date *time.Time) (*utils.Pagination, *response.RondaScheduleResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID, groupID, date)
	if err != nil {
		return nil, nil, err
	}

	resp := make(response.RondaScheduleResponses, len(data))
	for i, item := range data {
		totalMember, err := s.groupMemberRepository.GetTotalMember(item.GroupID)
		if err != nil {
			return nil, nil, err
		}
		resp[i] = *response.RondaScheduleModelToRondaScheduleResponse(&item, totalMember)
	}
	return paginated, &resp, err
}

func (s *rondaGroupScheduleServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
