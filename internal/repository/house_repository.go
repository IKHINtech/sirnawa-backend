package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type HouseRepository interface {
	Create(tx *gorm.DB, data models.House) (*models.House, error)
	Update(tx *gorm.DB, id string, data models.House) (*models.House, error)
	FindAll(rtID, blockID string) (models.Houses, error)
	Paginated(pagination utils.Pagination, rtID, blockID string) (*utils.Pagination, models.Houses, error)
	FindByID(id string) (*models.House, error)
	Delete(id string) error
}

type houseRepositoryImpl struct {
	db *gorm.DB
}

func NewHouseRepository(db *gorm.DB) HouseRepository {
	return &houseRepositoryImpl{db: db}
}

func (r *houseRepositoryImpl) Paginated(pagination utils.Pagination, rtID, blockID string) (*utils.Pagination, models.Houses, error) {
	var datas models.Houses
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}

	if blockID != "" {
		query = query.Where("block_id = ?", blockID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *houseRepositoryImpl) Create(tx *gorm.DB, data models.House) (*models.House, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *houseRepositoryImpl) Update(tx *gorm.DB, id string, data models.House) (*models.House, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.House{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *houseRepositoryImpl) FindByID(id string) (*models.House, error) {
	var data models.House

	query := r.db.
		Preload("Block").
		Preload("Rt").
		Preload("Rw").
		Preload("HousingArea")

	err := query.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *houseRepositoryImpl) FindAll(rtID, blockID string) (models.Houses, error) {
	var data models.Houses
	query := r.db
	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}

	if blockID != "" {
		query = query.Where("block_id = ?", blockID)
	}

	err := query.Find(&data).Error
	return data, err
}

func (r *houseRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.House{}, "id = ?", id).Error
	return err
}
