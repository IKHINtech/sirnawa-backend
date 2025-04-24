package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplService interface {
	Create(data request.IplCreateRequest) (*response.IplResponse, error)
	Update(id string, data request.IplUpdateRequset) (*response.IplResponse, error)
	FindByID(id string) (*response.IplFullResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.IplResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.IplResponses, error)
}

type iplServiceImpl struct {
	repository     repository.IplRepository
	itemRepository repository.IplDetailRepository
	db             *gorm.DB
}

func NewIplServices(
	repo repository.IplRepository,
	itemRepository repository.IplDetailRepository,
	db *gorm.DB,
) IplService {
	return &iplServiceImpl{repository: repo, itemRepository: itemRepository, db: db}
}

func (s *iplServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *iplServiceImpl) Create(data request.IplCreateRequest) (*response.IplResponse, error) {
	var result *models.Ipl

	err := s.withTransaction(func(tx *gorm.DB) error {
		// ubah semua yang aktif jadi tidak aktif
		err := s.repository.ChangeIsActiveByRtID(tx, data.RtID)
		if err != nil {
			return err
		}

		// buat payload
		payload := request.IplCreateRequestToIplModel(data)

		created, err := s.repository.Create(tx, payload)
		if err != nil {
			return err
		}

		// looping item
		for _, item := range data.Items {
			payloadItem := request.IplDetailCreateRequestToIplDetailModel(item)
			payloadItem.IplID = created.ID
			_, err = s.itemRepository.Create(tx, payloadItem)
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

	res := response.IplModelToIplResponse(result)
	return res, nil
}

func (s *iplServiceImpl) Update(id string, data request.IplUpdateRequset) (*response.IplResponse, error) {
	var result *models.Ipl

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.IplUpdateRequsetToIplModel(data)
		payload.ID = existing.ID

		updated, err := s.repository.Update(tx, payload.ID, payload)
		if err != nil {
			return err
		}

		// delete item sebelumnya
		err = s.itemRepository.DeleteByIplID(updated.ID)
		if err != nil {
			return err
		}

		// buat lagi item yang baru
		for _, item := range data.Items {
			payloadItem := request.IplDetailCreateRequestToIplDetailModel(item)
			payloadItem.IplID = updated.ID
			_, err = s.itemRepository.Create(tx, payloadItem)
			if err != nil {
				return err
			}
		}
		result = updated
		return nil
	})
	// handle error
	if err != nil {
		return nil, err
	}

	res := response.IplModelToIplResponse(result)
	return res, nil
}

func (s *iplServiceImpl) FindAll(rtID string) (response.IplResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.IplListToResponse(result)
	return resp, nil
}

func (s *iplServiceImpl) FindByID(id string) (*response.IplFullResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.IplModelToIplFullResponse(result)
	return resp, err
}

func (s *iplServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.IplResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.IplListToResponse(data)
	return paginated, &resp, err
}

func (s *iplServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
