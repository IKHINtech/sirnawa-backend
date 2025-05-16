package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplBillRepository interface {
	Create(tx *gorm.DB, data models.IplBill) (*models.IplBill, error)
	Update(tx *gorm.DB, id string, data models.IplBill) (*models.IplBill, error)
	FindAll(rtID, houseID, status string, month, year int) (models.IplBills, error)
	Paginated(pagination utils.Pagination, rtID, houseID, status string, month, year int) (*utils.Pagination, models.IplBills, error)
	FindByID(id string) (*models.IplBill, error)
	Delete(id string) error
}

type iplBillRepositoryImpl struct {
	db *gorm.DB
}

func NewIplBillRepository(db *gorm.DB) IplBillRepository {
	return &iplBillRepositoryImpl{db: db}
}

func (r *iplBillRepositoryImpl) Paginated(pagination utils.Pagination, rtID, houseID, status string, month, year int) (*utils.Pagination, models.IplBills, error) {
	var datas models.IplBills
	query := r.db.Preload("House").Preload("House.Block").Preload("Rt")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	if houseID != "" {
		query = query.Where("house_id = ?", houseID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if month != 0 {
		query = query.Where("month = ?", month)
	}
	if year != 0 {
		query = query.Where("year = ?", year)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplBillRepositoryImpl) Create(tx *gorm.DB, data models.IplBill) (*models.IplBill, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplBillRepositoryImpl) Update(tx *gorm.DB, id string, data models.IplBill) (*models.IplBill, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.IplBill{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplBillRepositoryImpl) FindByID(id string) (*models.IplBill, error) {
	var data models.IplBill

	err := r.db.Preload("House").Preload("House.Block").Preload("Rt").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplBillRepositoryImpl) FindAll(rtID, houseID, status string, month, year int) (models.IplBills, error) {
	query := r.db.Preload("House").Preload("House.Block").Preload("Rt")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}

	if houseID != "" {
		query = query.Where("house_id = ?", houseID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if month != 0 {
		query = query.Where("month = ?", month)
	}
	if year != 0 {
		query = query.Where("year = ?", year)
	}

	var data models.IplBills
	err := query.Find(&data).Error
	return data, err
}

func (r *iplBillRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.IplBill{}, "id = ?", id).Error
	return err
}
