package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ShopRepository interface {
	Create(tx *gorm.DB, data models.Shop) (*models.Shop, error)
	Update(tx *gorm.DB, id string, data models.Shop) (*models.Shop, error)
	FindAll() (models.Shops, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Shops, error)
	FindByID(id string) (*models.Shop, error)
	Delete(id string) error
}

type shopRepositoryImpl struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepositoryImpl{db: db}
}

func (r *shopRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Shops, error) {
	var datas models.Shops
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *shopRepositoryImpl) Create(tx *gorm.DB, data models.Shop) (*models.Shop, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *shopRepositoryImpl) Update(tx *gorm.DB, id string, data models.Shop) (*models.Shop, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Shop{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *shopRepositoryImpl) FindByID(id string) (*models.Shop, error) {
	var data models.Shop

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *shopRepositoryImpl) FindAll() (models.Shops, error) {
	var data models.Shops
	err := r.db.Find(&data).Error
	return data, err
}

func (r *shopRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Shop{}, "id = ?", id).Error
	return err
}
