package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplRateRepository interface {
	Create(tx *gorm.DB, data models.IplRate) (*models.IplRate, error)
	Update(tx *gorm.DB, id string, data models.IplRate) (*models.IplRate, error)
	FindAll(rtID string) (models.IplRates, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.IplRates, error)
	FindByID(id string) (*models.IplRate, error)
	Delete(id string) error
}

type iplRateRepositoryImpl struct {
	db *gorm.DB
}

func NewIplRateRepository(db *gorm.DB) IplRateRepository {
	return &iplRateRepositoryImpl{db: db}
}

func (r *iplRateRepositoryImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.IplRates, error) {
	var datas models.IplRates
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplRateRepositoryImpl) Create(tx *gorm.DB, data models.IplRate) (*models.IplRate, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplRateRepositoryImpl) Update(tx *gorm.DB, id string, data models.IplRate) (*models.IplRate, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.IplRate{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplRateRepositoryImpl) FindByID(id string) (*models.IplRate, error) {
	var data models.IplRate

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplRateRepositoryImpl) FindAll(rtID string) (models.IplRates, error) {
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	var data models.IplRates
	err := query.Find(&data).Error
	return data, err
}

func (r *iplRateRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.IplRate{}, "id = ?", id).Error
	return err
}
