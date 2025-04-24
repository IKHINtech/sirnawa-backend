package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type IplRepository interface {
	Create(tx *gorm.DB, data models.Ipl) (*models.Ipl, error)
	Update(tx *gorm.DB, id string, data models.Ipl) (*models.Ipl, error)
	FindAll() (models.Ipls, error)
	ChangeIsActiveByRtID(tx *gorm.DB, rtID string) error
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Ipls, error)
	FindByID(id string) (*models.Ipl, error)
	Delete(id string) error
}

type iplRepositoryImpl struct {
	db *gorm.DB
}

func NewIplRepository(db *gorm.DB) IplRepository {
	return &iplRepositoryImpl{db: db}
}

func (r *iplRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Ipls, error) {
	var datas models.Ipls
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *iplRepositoryImpl) Create(tx *gorm.DB, data models.Ipl) (*models.Ipl, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplRepositoryImpl) Update(tx *gorm.DB, id string, data models.Ipl) (*models.Ipl, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Ipl{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *iplRepositoryImpl) ChangeIsActiveByRtID(tx *gorm.DB, rtID string) error {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Ipl{}).Select("is_active").Where("rt_id = ?", rtID).Update("is_active", false).Error
	return err
}

func (r *iplRepositoryImpl) FindByID(id string) (*models.Ipl, error) {
	var data models.Ipl

	err := r.db.Preload("Items").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *iplRepositoryImpl) FindAll() (models.Ipls, error) {
	var data models.Ipls
	err := r.db.Find(&data).Error
	return data, err
}

func (r *iplRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Ipl{}, "id = ?", id).Error
	return err
}
