package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaContributionItemRepository interface {
	Create(tx *gorm.DB, data models.RondaContributionItem) (*models.RondaContributionItem, error)
	Update(tx *gorm.DB, id string, data models.RondaContributionItem) (*models.RondaContributionItem, error)
	FindAll() (models.RondaContributionItems, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaContributionItems, error)
	FindByID(id string) (*models.RondaContributionItem, error)
	Delete(id string) error
}

type rondaConstributionItemRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaContributionItemRepository(db *gorm.DB) RondaContributionItemRepository {
	return &rondaConstributionItemRepositoryImpl{db: db}
}

func (r *rondaConstributionItemRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaContributionItems, error) {
	var datas models.RondaContributionItems
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaConstributionItemRepositoryImpl) Create(tx *gorm.DB, data models.RondaContributionItem) (*models.RondaContributionItem, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaConstributionItemRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaContributionItem) (*models.RondaContributionItem, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaContributionItem{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaConstributionItemRepositoryImpl) FindByID(id string) (*models.RondaContributionItem, error) {
	var data models.RondaContributionItem

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaConstributionItemRepositoryImpl) FindAll() (models.RondaContributionItems, error) {
	var data models.RondaContributionItems
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rondaConstributionItemRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaContributionItem{}, "id = ?", id).Error
	return err
}
