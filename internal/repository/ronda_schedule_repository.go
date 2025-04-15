package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type BlockRepository interface {
	Create(tx *gorm.DB, data models.Block) (*models.Block, error)
	Update(tx *gorm.DB, id string, data models.Block) (*models.Block, error)
	FindAll() (models.Blocks, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Blocks, error)
	FindByID(id string) (*models.Block, error)
	Delete(id string) error
}

type blockRepositoryImpl struct {
	db *gorm.DB
}

func NewBlockRepository(db *gorm.DB) BlockRepository {
	return &blockRepositoryImpl{db: db}
}

func (r *blockRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Blocks, error) {
	var datas models.Blocks
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *blockRepositoryImpl) Create(tx *gorm.DB, data models.Block) (*models.Block, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *blockRepositoryImpl) Update(tx *gorm.DB, id string, data models.Block) (*models.Block, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Block{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *blockRepositoryImpl) FindByID(id string) (*models.Block, error) {
	var data models.Block

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *blockRepositoryImpl) FindAll() (models.Blocks, error) {
	var data models.Blocks
	err := r.db.Find(&data).Error
	return data, err
}

func (r *blockRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Block{}, "id = ?", id).Error
	return err
}
