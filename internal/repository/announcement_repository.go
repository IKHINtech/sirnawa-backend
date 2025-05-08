package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	Create(tx *gorm.DB, data models.Announcement) (*models.Announcement, error)
	Update(tx *gorm.DB, id string, data models.Announcement) (*models.Announcement, error)
	FindAll(rtID string) (models.Announcements, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.Announcements, error)
	FindByID(id string) (*models.Announcement, error)
	Delete(id string) error
}

type announcementRepositoryImpl struct {
	db *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) AnnouncementRepository {
	return &announcementRepositoryImpl{db: db}
}

func (r *announcementRepositoryImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, models.Announcements, error) {
	var datas models.Announcements
	query := r.db.Preload("User").Preload("User.Resident")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *announcementRepositoryImpl) Create(tx *gorm.DB, data models.Announcement) (*models.Announcement, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *announcementRepositoryImpl) Update(tx *gorm.DB, id string, data models.Announcement) (*models.Announcement, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Announcement{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *announcementRepositoryImpl) FindByID(id string) (*models.Announcement, error) {
	var data models.Announcement

	err := r.db.Preload("User").Preload("User.Resident").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *announcementRepositoryImpl) FindAll(rtID string) (models.Announcements, error) {
	var data models.Announcements

	query := r.db.Preload("User").Preload("User.Resident")
	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}
	err := query.Find(&data).Error
	return data, err
}

func (r *announcementRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Announcement{}, "id = ?", id).Error
	return err
}
