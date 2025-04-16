package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaAttendanceService interface {
	Create(data request.RondaAttendanceCreateRequest) (*response.RondaAttendanceResponse, error)
	Update(id string, data request.RondaAttendanceUpdateRequset) (*response.RondaAttendanceResponse, error)
	FindByID(id string) (*response.RondaAttendanceResponse, error)
	Delete(id string) error
	FindAll() (response.RondaAttendanceResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaAttendanceResponses, error)
}

type rondaAttendanceServiceImpl struct {
	repository repository.RondaAttendanceRepository
	db         *gorm.DB
}

func NewRondaAttendanceServices(repo repository.RondaAttendanceRepository, db *gorm.DB) RondaAttendanceService {
	return &rondaAttendanceServiceImpl{repository: repo, db: db}
}

func (s *rondaAttendanceServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *rondaAttendanceServiceImpl) Create(data request.RondaAttendanceCreateRequest) (*response.RondaAttendanceResponse, error) {
	var result *models.RondaAttendance

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.RondaAttendanceCreateRequestToRondaAttendanceModel(data)

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

	res := response.RondaAttendanceModelToRondaAttendanceResponse(result)
	return res, nil
}

func (s *rondaAttendanceServiceImpl) Update(id string, data request.RondaAttendanceUpdateRequset) (*response.RondaAttendanceResponse, error) {
	var result *models.RondaAttendance

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.RondaAttendanceUpdateRequsetToRondaAttendanceModel(data)
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

	res := response.RondaAttendanceModelToRondaAttendanceResponse(result)
	return res, nil
}

func (s *rondaAttendanceServiceImpl) FindAll() (response.RondaAttendanceResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.RondaAttendanceListToResponse(result)
	return resp, nil
}

func (s *rondaAttendanceServiceImpl) FindByID(id string) (*response.RondaAttendanceResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.RondaAttendanceModelToRondaAttendanceResponse(result)
	return resp, err
}

func (s *rondaAttendanceServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.RondaAttendanceResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.RondaAttendanceListToResponse(data)
	return paginated, &resp, err
}

func (s *rondaAttendanceServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
