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

type ResidentHouseService interface {
	AssignResidentToHouse(data request.ResidentHouseCreateRequest) (*response.ResidentHouseResponse, error)
	FindByHouseID(houseId string) ([]response.ResidentHouseDetailResponse, error)
	ChangeToPrimary(id string) error
	Delete(id string) error
}

type residentHouseServiceImpl struct {
	repository repository.ResidentHouseRepository
	db         *gorm.DB
}

func NewResidentHouseServices(repo repository.ResidentHouseRepository, db *gorm.DB) ResidentHouseService {
	return &residentHouseServiceImpl{repository: repo, db: db}
}

func (s *residentHouseServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *residentHouseServiceImpl) AssignResidentToHouse(data request.ResidentHouseCreateRequest) (*response.ResidentHouseResponse, error) {
	var result *models.ResidentHouse

	err := s.withTransaction(func(tx *gorm.DB) error {
		allHouses, err := s.repository.FindByResidentID(data.ResidentID)
		if err != nil {
			return err
		} // Cek duplikat HouseID
		for _, house := range allHouses {
			if house.HouseID == data.HouseID {
				return errors.New("rumah ini sudah terdaftar untuk penghuni tersebut")
			}
		}
		payload := request.ResidentHouseCreateRequestToResidentHouseModel(data)

		if payload.IsPrimary {
			// ubah rumah yang lain jadi not primary
			err := s.repository.ChangeNotPrimaryByResidentID(tx, data.ResidentID)
			if err != nil {
				return err
			}
		}

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

	res := response.ResidentHouseModelToResidentHouseResponse(result)
	return res, nil
}

func (s *residentHouseServiceImpl) Update(id string, data request.ResidentHouseUpdateRequset) (*response.ResidentHouseResponse, error) {
	var result *models.ResidentHouse

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.ResidentHouseUpdateRequsetToResidentHouseModel(data)
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

	res := response.ResidentHouseModelToResidentHouseResponse(result)
	return res, nil
}

func (s *residentHouseServiceImpl) FindByHouseID(houseID string) ([]response.ResidentHouseDetailResponse, error) {
	result, err := s.repository.FindAll(houseID)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return []response.ResidentHouseDetailResponse{}, nil
	}

	res := make([]response.ResidentHouseDetailResponse, len(result))
	for i, v := range result {
		res[i] = *response.MapResidentHouseDetailResponse(&v)
	}

	return res, nil
}

func (s *residentHouseServiceImpl) FindAll(houseID string) (response.ResidentHouseResponses, error) {
	result, err := s.repository.FindAll(houseID)
	if err != nil {
		return nil, err
	}

	resp := response.ResidentHouseListToResponse(result)
	return resp, nil
}

func (s *residentHouseServiceImpl) FindByID(id string) (*response.ResidentHouseResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.ResidentHouseModelToResidentHouseResponse(result)
	return resp, err
}

func (s *residentHouseServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.ResidentHouseResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.ResidentHouseListToResponse(data)
	return paginated, &resp, err
}

func (s *residentHouseServiceImpl) ChangeToPrimary(id string) error {
	return s.withTransaction(func(tx *gorm.DB) error {
		existingData, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		err = s.repository.ChangeNotPrimaryByResidentID(tx, existingData.ResidentID)
		if err != nil {
			return err
		}

		err = s.repository.ChangeToPrimary(tx, existingData.ID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *residentHouseServiceImpl) Delete(id string) error {
	return s.withTransaction(func(tx *gorm.DB) error {
		existingData, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		allHouses, err := s.repository.FindByResidentID(existingData.ResidentID)
		if err != nil {
			return err
		}

		if len(allHouses) == 1 && allHouses[0].ID == existingData.ID {
			// Jika rumah hanya 1, tidak boleh menghapus rumah
			return errors.New("tidak bisa menghapus rumah utama")
		}

		if existingData.IsPrimary {
			// Karena primary yang dihapus, harus ada rumah lain yang dijadikan primary
			for _, house := range allHouses {
				if house.ID != existingData.ID {
					err = s.repository.ChangeToPrimary(tx, house.ID)
					if err != nil {
						return err
					}
					break // cukup ubah satu rumah jadi primary
				}
			}
		}

		// Hapus rumah
		err = s.repository.Delete(tx, id)
		if err != nil {
			return err
		}

		return nil
	})
}
