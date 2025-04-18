package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RtRepository interface {
	Create(tx *gorm.DB, data models.Rt) (*models.Rt, error)
	Update(tx *gorm.DB, id string, data models.Rt) (*models.Rt, error)
	FindAll() (models.Rts, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Rts, error)
	FindByID(id string) (*models.Rt, error)
	Delete(id string) error
}

type rtRepositoryImpl struct {
	db *gorm.DB
}

func NewRtRepository(db *gorm.DB) RtRepository {
	return &rtRepositoryImpl{db: db}
}

func (r *rtRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Rts, error) {
	var datas models.Rts
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rtRepositoryImpl) Create(tx *gorm.DB, data models.Rt) (*models.Rt, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rtRepositoryImpl) Update(tx *gorm.DB, id string, data models.Rt) (*models.Rt, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Rt{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rtRepositoryImpl) FindByID(id string) (*models.Rt, error) {
	var data models.Rt

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rtRepositoryImpl) FindAll() (models.Rts, error) {
	var data models.Rts
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rtRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Rt{}, "id = ?", id).Error
	return err
}
