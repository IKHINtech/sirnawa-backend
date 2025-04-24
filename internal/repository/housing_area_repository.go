package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type HousingAreaRepository interface {
	Create(tx *gorm.DB, data models.HousingArea) (*models.HousingArea, error)
	Update(tx *gorm.DB, id string, data models.HousingArea) (*models.HousingArea, error)
	FindAll() (models.HousingAreas, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.HousingAreas, error)
	FindByID(id string) (*models.HousingArea, error)
	Delete(id string) error
}

type housingAreaRepositoryImpl struct {
	db *gorm.DB
}

func NewHousingAreaRepository(db *gorm.DB) HousingAreaRepository {
	return &housingAreaRepositoryImpl{db: db}
}

func (r *housingAreaRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.HousingAreas, error) {
	var datas models.HousingAreas
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *housingAreaRepositoryImpl) Create(tx *gorm.DB, data models.HousingArea) (*models.HousingArea, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *housingAreaRepositoryImpl) Update(tx *gorm.DB, id string, data models.HousingArea) (*models.HousingArea, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.HousingArea{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *housingAreaRepositoryImpl) FindByID(id string) (*models.HousingArea, error) {
	var data models.HousingArea

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *housingAreaRepositoryImpl) FindAll() (models.HousingAreas, error) {
	var data models.HousingAreas
	err := r.db.Find(&data).Error
	return data, err
}

func (r *housingAreaRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.HousingArea{}, "id = ?", id).Error
	return err
}
