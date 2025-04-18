package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RwRepository interface {
	Create(tx *gorm.DB, data models.Rw) (*models.Rw, error)
	Update(tx *gorm.DB, id string, data models.Rw) (*models.Rw, error)
	FindAll() (models.Rws, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Rws, error)
	FindByID(id string) (*models.Rw, error)
	Delete(id string) error
}

type rwRepositoryImpl struct {
	db *gorm.DB
}

func NewRwRepository(db *gorm.DB) RwRepository {
	return &rwRepositoryImpl{db: db}
}

func (r *rwRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Rws, error) {
	var datas models.Rws
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rwRepositoryImpl) Create(tx *gorm.DB, data models.Rw) (*models.Rw, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rwRepositoryImpl) Update(tx *gorm.DB, id string, data models.Rw) (*models.Rw, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Rw{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rwRepositoryImpl) FindByID(id string) (*models.Rw, error) {
	var data models.Rw

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rwRepositoryImpl) FindAll() (models.Rws, error) {
	var data models.Rws
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rwRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Rw{}, "id = ?", id).Error
	return err
}
