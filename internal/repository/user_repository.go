package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, data models.User) (*models.User, error)
	Update(tx *gorm.DB, id string, data models.User) (*models.User, error)
	FindAll() (models.Users, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Users, error)
	FindByID(id string) (*models.User, error)
	Delete(id string) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Users, error) {
	var datas models.Users
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *userRepositoryImpl) Create(tx *gorm.DB, data models.User) (*models.User, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *userRepositoryImpl) Update(tx *gorm.DB, id string, data models.User) (*models.User, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *userRepositoryImpl) FindByID(id string) (*models.User, error) {
	var data models.User

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *userRepositoryImpl) FindAll() (models.Users, error) {
	var data models.Users
	err := r.db.Find(&data).Error
	return data, err
}

func (r *userRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.User{}, "id = ?", id).Error
	return err
}
