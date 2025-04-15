package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ShopProductRepository interface {
	Create(tx *gorm.DB, data models.ShopProduct) (*models.ShopProduct, error)
	Update(tx *gorm.DB, id string, data models.ShopProduct) (*models.ShopProduct, error)
	FindAll() (models.ShopProducts, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.ShopProducts, error)
	FindByID(id string) (*models.ShopProduct, error)
	Delete(id string) error
}

type shopProductRepositoryImpl struct {
	db *gorm.DB
}

func NewShopProductRepository(db *gorm.DB) ShopProductRepository {
	return &shopProductRepositoryImpl{db: db}
}

func (r *shopProductRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.ShopProducts, error) {
	var datas models.ShopProducts
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *shopProductRepositoryImpl) Create(tx *gorm.DB, data models.ShopProduct) (*models.ShopProduct, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *shopProductRepositoryImpl) Update(tx *gorm.DB, id string, data models.ShopProduct) (*models.ShopProduct, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.ShopProduct{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *shopProductRepositoryImpl) FindByID(id string) (*models.ShopProduct, error) {
	var data models.ShopProduct

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *shopProductRepositoryImpl) FindAll() (models.ShopProducts, error) {
	var data models.ShopProducts
	err := r.db.Find(&data).Error
	return data, err
}

func (r *shopProductRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.ShopProduct{}, "id = ?", id).Error
	return err
}
