package services

import (
	"errors"

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
	FindByID(id string) (*response.HouseResponseDetail, error)
	Delete(id string) error
	FindAll(rtID string) (response.HouseResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.HouseResponses, error)
}

type houseServiceImpl struct {
	repository repository.HouseRepository
	rwRepo     repository.RwRepository
	db         *gorm.DB
}

func NewHouseServices(
	repo repository.HouseRepository,
	rwRepo repository.RwRepository,
	db *gorm.DB,
) HouseService {
	return &houseServiceImpl{
		repository: repo,
		rwRepo:     rwRepo,
		db:         db,
	}
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
		rw, err := s.rwRepo.FindByRtID(data.RtID)
		if err != nil {
			return err
		}

		if rw == nil {
			return errors.New("rw not found")
		}

		payload := request.HouseCreateRequestToHouseModel(data, rw.ID, rw.HousingAreaID)

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

		payload := request.HouseUpdateRequsetToHouseModel(data, existing.RwID, existing.HousingAreaID)
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

func (s *houseServiceImpl) FindAll(rtID string) (response.HouseResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.HouseListToResponse(result)
	return resp, nil
}

func (s *houseServiceImpl) FindByID(id string) (*response.HouseResponseDetail, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.MapHouseDetailResponse(result)
	return resp, err
}

func (s *houseServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.HouseResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
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
