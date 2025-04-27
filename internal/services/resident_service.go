package services

import (
	"errors"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ResidentService interface {
	Create(data request.ResidentCreateRequest) (*response.ResidentResponse, error)
	Update(id string, data request.ResidentUpdateRequset) (*response.ResidentResponse, error)
	FindByID(id string) (*response.ResidentResponse, error)
	Delete(id string) error
	FindAll() (response.ResidentResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ResidentResponses, error)
}

type residentServiceImpl struct {
	repository      repository.ResidentRepository
	userRepostitory repository.UserRepository
	db              *gorm.DB
}

func NewResidentServices(
	repo repository.ResidentRepository,
	userRepostitory repository.UserRepository,
	db *gorm.DB,
) ResidentService {
	return &residentServiceImpl{
		repository:      repo,
		userRepostitory: userRepostitory,
		db:              db,
	}
}

func (s *residentServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *residentServiceImpl) Create(data request.ResidentCreateRequest) (*response.ResidentResponse, error) {
	var result *models.Resident

	err := s.withTransaction(func(tx *gorm.DB) error {
		// cari dulu apakah nik sudah terdaftar
		existingNik, err := s.repository.FindByNIK(data.NIK)
		if err != nil {
			return err
		}
		if existingNik != nil {
			// kalau sudah ter register maka return
			return errors.New("NIK already registered")
		}

		payload := request.ResidentCreateRequestToResidentModel(data)

		created, err := s.repository.Create(tx, payload)
		if err != nil {
			return err
		}

		if data.Email != nil {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.AppConfig.DEFAULT_PASSWORD), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			// ada data email pada resident request, buatkan user
			payloadUser := models.User{
				Email:      *data.Email,
				Password:   string(hashedPassword),
				Role:       models.Role(data.Role),
				ResidentID: &created.ID,
			}

			_, err = s.userRepostitory.Create(tx, payloadUser)
			if err != nil {
				return err
			}
		}

		result = created
		return nil
	})
	// handling err
	if err != nil {
		return nil, err
	}

	res := response.ResidentModelToResidentResponse(result)
	return res, nil
}

func (s *residentServiceImpl) Update(id string, data request.ResidentUpdateRequset) (*response.ResidentResponse, error) {
	var result *models.Resident

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.ResidentUpdateRequsetToResidentModel(data)
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

	res := response.ResidentModelToResidentResponse(result)
	return res, nil
}

func (s *residentServiceImpl) FindAll() (response.ResidentResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.ResidentListToResponse(result)
	return resp, nil
}

func (s *residentServiceImpl) FindByID(id string) (*response.ResidentResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.ResidentModelToResidentResponse(result)
	return resp, err
}

func (s *residentServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ResidentResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.ResidentListToResponse(data)
	return paginated, &resp, err
}

func (s *residentServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
