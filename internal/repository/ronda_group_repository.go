package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaGroupRepository interface {
	Create(tx *gorm.DB, data models.RondaGroup) (*models.RondaGroup, error)
	Update(tx *gorm.DB, id string, data models.RondaGroup) (*models.RondaGroup, error)
	FindAll(rtID string) (models.RondaGroups, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.RondaGroups, error)
	FindByID(id string) (*models.RondaGroup, error)
	Delete(id string) error
}

type rondaGroupRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaGroupRepository(db *gorm.DB) RondaGroupRepository {
	return &rondaGroupRepositoryImpl{db: db}
}

func (r *rondaGroupRepositoryImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.RondaGroups, error) {
	var datas models.RondaGroups
	query := r.db.Preload("Rt")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaGroupRepositoryImpl) Create(tx *gorm.DB, data models.RondaGroup) (*models.RondaGroup, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaGroupRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaGroup) (*models.RondaGroup, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaGroup{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaGroupRepositoryImpl) FindByID(id string) (*models.RondaGroup, error) {
	var data models.RondaGroup

	err := r.db.Preload("Rt").Preload("Members").Preload("Members.Resident").Preload("Members.House").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaGroupRepositoryImpl) FindAll(rtID string) (models.RondaGroups, error) {
	var data models.RondaGroups
	query := r.db.Preload("Rt")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Find(&data).Error
	return data, err
}

func (r *rondaGroupRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaGroup{}, "id = ?", id).Error
	return err
}
