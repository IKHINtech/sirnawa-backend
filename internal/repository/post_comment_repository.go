package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type PostCommentRepository interface {
	Create(tx *gorm.DB, data models.PostComment) (*models.PostComment, error)
	Update(tx *gorm.DB, id string, data models.PostComment) (*models.PostComment, error)
	FindAll() (models.PostComments, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.PostComments, error)
	FindByID(id string) (*models.PostComment, error)
	Delete(id string) error
}

type postCommentRepositoryImpl struct {
	db *gorm.DB
}

func NewPostCommentRepository(db *gorm.DB) PostCommentRepository {
	return &postCommentRepositoryImpl{db: db}
}

func (r *postCommentRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.PostComments, error) {
	var datas models.PostComments
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *postCommentRepositoryImpl) Create(tx *gorm.DB, data models.PostComment) (*models.PostComment, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *postCommentRepositoryImpl) Update(tx *gorm.DB, id string, data models.PostComment) (*models.PostComment, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.PostComment{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *postCommentRepositoryImpl) FindByID(id string) (*models.PostComment, error) {
	var data models.PostComment

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *postCommentRepositoryImpl) FindAll() (models.PostComments, error) {
	var data models.PostComments
	err := r.db.Find(&data).Error
	return data, err
}

func (r *postCommentRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.PostComment{}, "id = ?", id).Error
	return err
}
