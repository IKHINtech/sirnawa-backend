package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(tx *gorm.DB, data models.Item) (*models.Item, error)
	Update(tx *gorm.DB, id string, data models.Item) (*models.Item, error)
	FindAll(rtID string) (models.Items, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.Items, error)
	FindByID(id string) (*models.Item, error)
	Delete(id string) error
}

type itemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepositoryImpl{db: db}
}

func (r *itemRepositoryImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.Items, error) {
	var datas models.Items
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *itemRepositoryImpl) Create(tx *gorm.DB, data models.Item) (*models.Item, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *itemRepositoryImpl) Update(tx *gorm.DB, id string, data models.Item) (*models.Item, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Item{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *itemRepositoryImpl) FindByID(id string) (*models.Item, error) {
	var data models.Item

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *itemRepositoryImpl) FindAll(rtID string) (models.Items, error) {
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	var data models.Items
	err := query.Find(&data).Error
	return data, err
}

func (r *itemRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Item{}, "id = ?", id).Error
	return err
}
