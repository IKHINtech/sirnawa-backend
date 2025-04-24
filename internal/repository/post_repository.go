package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(tx *gorm.DB, data models.Post) (*models.Post, error)
	Update(tx *gorm.DB, id string, data models.Post) (*models.Post, error)
	FindAll(rtID string) (models.Posts, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.Posts, error)
	FindByID(id string) (*models.Post, error)
	Delete(id string) error
}

type postRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.Posts, error) {
	var datas models.Posts
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *postRepositoryImpl) Create(tx *gorm.DB, data models.Post) (*models.Post, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *postRepositoryImpl) Update(tx *gorm.DB, id string, data models.Post) (*models.Post, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Post{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *postRepositoryImpl) FindByID(id string) (*models.Post, error) {
	var data models.Post

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *postRepositoryImpl) FindAll(rtID string) (models.Posts, error) {
	var data models.Posts
	query := r.db

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Find(&data).Error
	return data, err
}

func (r *postRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Post{}, "id = ?", id).Error
	return err
}
