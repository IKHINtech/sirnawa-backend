package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplPaymentRepository interface {
	Create(tx *gorm.DB, data models.IplPayment) (*models.IplPayment, error)
	Update(tx *gorm.DB, id string, data models.IplPayment) (*models.IplPayment, error)
	FindAll() (models.IplPayments, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.IplPayments, error)
	FindByID(id string) (*models.IplPayment, error)
	Delete(id string) error
}

type iplPaymnetRepositoryImpl struct {
	db *gorm.DB
}

func NewIplPaymentRepository(db *gorm.DB) IplPaymentRepository {
	return &iplPaymnetRepositoryImpl{db: db}
}

func (r *iplPaymnetRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.IplPayments, error) {
	var datas models.IplPayments
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplPaymnetRepositoryImpl) Create(tx *gorm.DB, data models.IplPayment) (*models.IplPayment, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplPaymnetRepositoryImpl) Update(tx *gorm.DB, id string, data models.IplPayment) (*models.IplPayment, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.IplPayment{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplPaymnetRepositoryImpl) FindByID(id string) (*models.IplPayment, error) {
	var data models.IplPayment

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplPaymnetRepositoryImpl) FindAll() (models.IplPayments, error) {
	var data models.IplPayments
	err := r.db.Find(&data).Error
	return data, err
}

func (r *iplPaymnetRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.IplPayment{}, "id = ?", id).Error
	return err
}
