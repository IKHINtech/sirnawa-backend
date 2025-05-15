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

type IplBillService interface {
	Create(data request.IplBillCreateRequest) (*response.IplBillResponse, error)
	Update(id string, data request.IplBillUpdateRequset) (*response.IplBillResponse, error)
	FindByID(id string) (*response.IplBillResponse, error)
	Delete(id string) error
	GenerateIplBill(data request.IplBillGenerator) error
	FindAll(rtID, houseID, status string, month, years int) (response.IplBillResponses, error)
	Paginated(pagination utils.Pagination, rtID, houseID, status string, month, year int) (*utils.Pagination, *response.IplBillResponses, error)
}

type iplBillServiceImpl struct {
	repository        repository.IplBillRepository
	iplRateRepo       repository.IplRateRepository
	iplRateDetailRepo repository.IplRateDetailRepository
	iplBillDetailRepo repository.IplBillDetailRepository
	houseRepo         repository.HouseRepository
	db                *gorm.DB
}

func NewIplBillServices(
	repo repository.IplBillRepository,
	iplRateRepo repository.IplRateRepository,
	iplRateDetailRepo repository.IplRateDetailRepository,
	iplBillDetailRepo repository.IplBillDetailRepository,
	houseRepo repository.HouseRepository,
	db *gorm.DB,
) IplBillService {
	return &iplBillServiceImpl{
		repository:        repo,
		iplRateRepo:       iplRateRepo,
		iplRateDetailRepo: iplRateDetailRepo,
		iplBillDetailRepo: iplBillDetailRepo,
		db:                db,
	}
}

func (s *iplBillServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *iplBillServiceImpl) Create(data request.IplBillCreateRequest) (*response.IplBillResponse, error) {
	var result *models.IplBill

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.IplBillCreateRequestToIplBillModel(data)

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

	res := response.IplBillModelToIplBillResponse(result)
	return res, nil
}

func (s *iplBillServiceImpl) Generate(tx *gorm.DB, data request.IplBillGenerator, house models.House, iplRate models.IplRate, iplRateDetails []models.IplRateDetail) error {
	var nol *int64
	payload := models.IplBill{
		HouseID:     house.ID,
		RtID:        data.RtID,
		Year:        data.Year,
		Month:       data.Month,
		TotalAmount: iplRate.Amount,
		BalanceDue:  nol,
		AmountPaid:  nol,
		IplRateID:   &data.IplRateID,
		Status:      models.IplBillStatusUnpaid,
		DueDate:     time.Date(data.Year, time.Month(data.Month), data.DueDate, 0, 0, 0, 0, time.UTC),
		Penalty:     nil,
	}
	newBill, err := s.repository.Create(tx, payload)
	if err != nil {
		return err
	}
	// buat detail bill dari detail ipl rate
	for _, detail := range iplRateDetails {
		payloadDetail := models.IplBillDetail{
			IplBillID: newBill.ID,
			ItemID:    detail.ItemID,
			SubAmount: detail.Amount,
		}
		_, err := s.iplBillDetailRepo.Create(tx, payloadDetail)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *iplBillServiceImpl) GenerateIplBill(data request.IplBillGenerator) error {
	err := s.withTransaction(func(tx *gorm.DB) error {
		// cari ipl rate
		iplRate, err := s.iplRateRepo.FindByID(data.IplRateID)
		if err != nil {
			return err
		}

		iplRateDetails, err := s.iplRateDetailRepo.FindAll(data.IplRateID)
		if err != nil {
			return err
		}

		status := "tidak aktif"
		if data.IsAllHouse {
			allHouse, err := s.houseRepo.FindAll(data.RtID, "", "", "", &status)
			if err != nil {
				return err
			}
			for _, house := range allHouse {
				err := s.Generate(tx, data, house, *iplRate, iplRateDetails)
				if err != nil {
					return err
				}
			}
		} else {
			houses, err := s.houseRepo.FindByIDs(data.HouseIDs)
			if err != nil {
				return err
			}
			for _, house := range houses {
				err := s.Generate(tx, data, house, *iplRate, iplRateDetails)
				if err != nil {
					return err
				}
			}
		}

		return err
	})
	return err
}

func (s *iplBillServiceImpl) Update(id string, data request.IplBillUpdateRequset) (*response.IplBillResponse, error) {
	var result *models.IplBill

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.IplBillUpdateRequsetToIplBillModel(data)
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

	res := response.IplBillModelToIplBillResponse(result)
	return res, nil
}

func (s *iplBillServiceImpl) FindAll(rtID, houseID, status string, month, year int) (response.IplBillResponses, error) {
	result, err := s.repository.FindAll(rtID, houseID, status, month, year)
	if err != nil {
		return nil, err
	}

	resp := response.IplBillListToResponse(result)
	return resp, nil
}

func (s *iplBillServiceImpl) FindByID(id string) (*response.IplBillResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.IplBillModelToIplBillResponse(result)
	return resp, err
}

func (s *iplBillServiceImpl) Paginated(pagination utils.Pagination, rtID, houseID, status string, month, year int) (*utils.Pagination, *response.IplBillResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID, houseID, status, month, year)
	if err != nil {
		return nil, nil, err
	}

	resp := response.IplBillListToResponse(data)
	return paginated, &resp, err
}

func (s *iplBillServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
