package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type UserRTRepository interface {
	Create(tx *gorm.DB, data models.UserRT) (*models.UserRT, error)
	Update(tx *gorm.DB, id string, data models.UserRT) (*models.UserRT, error)
	FindAll(rtID string) (models.UserRTs, error)
	FindByUserID(userID string) (models.UserRTs, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.UserRTs, error)
	FindByID(id string) (*models.UserRT, error)
	Delete(id string) error
}

type userRTRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRTRepository(db *gorm.DB) UserRTRepository {
	return &userRTRepositoryImpl{db: db}
}

func (r *userRTRepositoryImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.UserRTs, error) {
	var datas models.UserRTs
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *userRTRepositoryImpl) Create(tx *gorm.DB, data models.UserRT) (*models.UserRT, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *userRTRepositoryImpl) Update(tx *gorm.DB, id string, data models.UserRT) (*models.UserRT, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.UserRT{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *userRTRepositoryImpl) FindByID(id string) (*models.UserRT, error) {
	var data models.UserRT

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *userRTRepositoryImpl) FindByUserID(userID string) (models.UserRTs, error) {
	query := r.db.Where("user_id = ?", userID)
	var data models.UserRTs
	err := query.Find(&data).Error
	return data, err
}

func (r *userRTRepositoryImpl) FindAll(rtID string) (models.UserRTs, error) {
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	var data models.UserRTs
	err := query.Find(&data).Error
	return data, err
}

func (r *userRTRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.UserRT{}, "id = ?", id).Error
	return err
}
