package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaConstributionRepository interface {
	Create(tx *gorm.DB, data models.RondaConstribution) (*models.RondaConstribution, error)
	Update(tx *gorm.DB, id string, data models.RondaConstribution) (*models.RondaConstribution, error)
	FindAll() (models.RondaConstributions, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaConstributions, error)
	FindByID(id string) (*models.RondaConstribution, error)
	Delete(id string) error
}

type rondaConstributionRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaConstributionRepository(db *gorm.DB) RondaConstributionRepository {
	return &rondaConstributionRepositoryImpl{db: db}
}

func (r *rondaConstributionRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaConstributions, error) {
	var datas models.RondaConstributions
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaConstributionRepositoryImpl) Create(tx *gorm.DB, data models.RondaConstribution) (*models.RondaConstribution, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaConstributionRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaConstribution) (*models.RondaConstribution, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaConstribution{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaConstributionRepositoryImpl) FindByID(id string) (*models.RondaConstribution, error) {
	var data models.RondaConstribution

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaConstributionRepositoryImpl) FindAll() (models.RondaConstributions, error) {
	var data models.RondaConstributions
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rondaConstributionRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaConstribution{}, "id = ?", id).Error
	return err
}
