package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type BaseRepository interface {
	Create(tx *gorm.DB, data models.BaseModel) (*models.BaseModel, error)
	Update(tx *gorm.DB, id string, data models.BaseModel) (*models.BaseModel, error)
	FindAll() (models.BaseModels, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.BaseModels, error)
	FindByID(id string) (*models.BaseModel, error)
	Delete(id string) error
}

type BaseRepositoryImpl struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &BaseRepositoryImpl{db: db}
}

func (r *BaseRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.BaseModels, error) {
	var datas models.BaseModels
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *BaseRepositoryImpl) Create(tx *gorm.DB, data models.BaseModel) (*models.BaseModel, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *BaseRepositoryImpl) Update(tx *gorm.DB, id string, data models.BaseModel) (*models.BaseModel, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.BaseModel{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *BaseRepositoryImpl) FindByID(id string) (*models.BaseModel, error) {
	var data models.BaseModel

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *BaseRepositoryImpl) FindAll() (models.BaseModels, error) {
	var data models.BaseModels
	err := r.db.Find(&data).Error
	return data, err
}

func (r *BaseRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.BaseModel{}, "id = ?", id).Error
	return err
}
