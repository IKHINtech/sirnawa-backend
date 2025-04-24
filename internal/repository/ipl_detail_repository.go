package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplDetailRepository interface {
	Create(tx *gorm.DB, data models.IplDetail) (*models.IplDetail, error)
	Update(tx *gorm.DB, id string, data models.IplDetail) (*models.IplDetail, error)
	FindAll() (models.IplDetails, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.IplDetails, error)
	FindByID(id string) (*models.IplDetail, error)
	Delete(id string) error
	DeleteByIplID(iplID string) error
}

type iplDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewIplDetailRepository(db *gorm.DB) IplDetailRepository {
	return &iplDetailRepositoryImpl{db: db}
}

func (r *iplDetailRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.IplDetails, error) {
	var datas models.IplDetails
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplDetailRepositoryImpl) Create(tx *gorm.DB, data models.IplDetail) (*models.IplDetail, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplDetailRepositoryImpl) Update(tx *gorm.DB, id string, data models.IplDetail) (*models.IplDetail, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.IplDetail{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplDetailRepositoryImpl) FindByID(id string) (*models.IplDetail, error) {
	var data models.IplDetail

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplDetailRepositoryImpl) FindAll() (models.IplDetails, error) {
	var data models.IplDetails
	err := r.db.Find(&data).Error
	return data, err
}

func (r *iplDetailRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.IplDetail{}, "id = ?", id).Error
	return err
}

func (r *iplDetailRepositoryImpl) DeleteByIplID(iplID string) error {
	err := r.db.Delete(&models.IplDetail{}, "ipl_id = ?", iplID).Error
	return err
}
