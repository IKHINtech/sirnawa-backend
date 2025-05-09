package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplRateDetailRepository interface {
	Create(tx *gorm.DB, data models.IplRateDetail) (*models.IplRateDetail, error)
	Update(tx *gorm.DB, id string, data models.IplRateDetail) (*models.IplRateDetail, error)
	FindAll(iplRateID string) (models.IplRateDetails, error)
	Paginated(pagination utils.Pagination, iplRateID string) (*utils.Pagination, models.IplRateDetails, error)
	FindByID(id string) (*models.IplRateDetail, error)
	Delete(id string) error
}

type iplRateDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewIplRateDetailRepository(db *gorm.DB) IplRateDetailRepository {
	return &iplRateDetailRepositoryImpl{db: db}
}

func (r *iplRateDetailRepositoryImpl) Paginated(pagination utils.Pagination, iplRateID string) (*utils.Pagination, models.IplRateDetails, error) {
	var datas models.IplRateDetails
	query := r.db

	if iplRateID != "" {
		query = query.Where("ipl_rate_id = ?", iplRateID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplRateDetailRepositoryImpl) Create(tx *gorm.DB, data models.IplRateDetail) (*models.IplRateDetail, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplRateDetailRepositoryImpl) Update(tx *gorm.DB, id string, data models.IplRateDetail) (*models.IplRateDetail, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.IplRateDetail{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplRateDetailRepositoryImpl) FindByID(id string) (*models.IplRateDetail, error) {
	var data models.IplRateDetail

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplRateDetailRepositoryImpl) FindAll(iplRateID string) (models.IplRateDetails, error) {
	query := r.db

	if iplRateID != "" {
		query = query.Where("ipl_rate_id = ?", iplRateID)
	}
	var data models.IplRateDetails
	err := query.Find(&data).Error
	return data, err
}

func (r *iplRateDetailRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.IplRateDetail{}, "id = ?", id).Error
	return err
}
