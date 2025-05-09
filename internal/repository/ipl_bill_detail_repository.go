package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplBillDetailRepository interface {
	Create(tx *gorm.DB, data models.IplBillDetail) (*models.IplBillDetail, error)
	Update(tx *gorm.DB, id string, data models.IplBillDetail) (*models.IplBillDetail, error)
	FindAll(iplBillDetail string) (models.IplBillDetails, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.IplBillDetails, error)
	FindByID(id string) (*models.IplBillDetail, error)
	Delete(id string) error
}

type iplBillDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewIplBillDetailRepository(db *gorm.DB) IplBillDetailRepository {
	return &iplBillDetailRepositoryImpl{db: db}
}

func (r *iplBillDetailRepositoryImpl) Paginated(pagination utils.Pagination, iplBillID string) (*utils.Pagination, models.IplBillDetails, error) {
	var datas models.IplBillDetails
	query := r.db

	if iplBillID != "" {
		query = query.Where("ipl_bill_id = ?", iplBillID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplBillDetailRepositoryImpl) Create(tx *gorm.DB, data models.IplBillDetail) (*models.IplBillDetail, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplBillDetailRepositoryImpl) Update(tx *gorm.DB, id string, data models.IplBillDetail) (*models.IplBillDetail, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.IplBillDetail{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplBillDetailRepositoryImpl) FindByID(id string) (*models.IplBillDetail, error) {
	var data models.IplBillDetail

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplBillDetailRepositoryImpl) FindAll(iplBillId string) (models.IplBillDetails, error) {
	query := r.db

	if iplBillId != "" {
		query = query.Where("ipl_bill_id = ?", iplBillId)
	}
	var data models.IplBillDetails
	err := query.Find(&data).Error
	return data, err
}

func (r *iplBillDetailRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.IplBillDetail{}, "id = ?", id).Error
	return err
}
