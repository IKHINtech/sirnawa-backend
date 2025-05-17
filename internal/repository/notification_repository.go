package repository

import (
	"context"
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	GetByID(ctx context.Context, id string) (*models.Notification, error)
	GetByUserID(pagination utils.Pagination, userID, rtID, houseID string) (*utils.Pagination, []models.Notification, error)
	MarkAsRead(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	Delete(ctx context.Context, id string) error
	CountUnread(ctx context.Context, userID string) (int64, error)
}

type notificationRepositoryImpl struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepositoryImpl{db: db}
}

// Create membuat notifikasi baru
func (r *notificationRepositoryImpl) Create(ctx context.Context, notification *models.Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

// GetByID mendapatkan notifikasi berdasarkan ID
func (r *notificationRepositoryImpl) GetByID(ctx context.Context, id string) (*models.Notification, error) {
	var notification models.Notification
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&notification).Error
	return &notification, err
}

// GetByUserID mendapatkan notifikasi untuk user tertentu
func (r *notificationRepositoryImpl) GetByUserID(pagination utils.Pagination, userID, rtID, houseID string) (*utils.Pagination, []models.Notification, error) {
	var notifications []models.Notification
	query := r.db.Where("user_id")

	if houseID != "" {
		query = query.Where("house_id", houseID)
	}

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}

	err := query.Scopes(utils.Paginate(notifications, &pagination, query)).Find(&notifications).Error

	return &pagination, notifications, err
}

// MarkAsRead menandai notifikasi sebagai sudah dibaca
func (r *notificationRepositoryImpl) MarkAsRead(ctx context.Context, id string) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("id = ?", id).
		Update("read_at", &now).Error
}

// MarkAllAsRead menandai semua notifikasi user sebagai sudah dibaca
func (r *notificationRepositoryImpl) MarkAllAsRead(ctx context.Context, userID string) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("user_id = ? AND read_at IS NULL", userID).
		Update("read_at", &now).Error
}

// Delete menghapus notifikasi
func (r *notificationRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Notification{}, "id = ?", id).Error
}

// CountUnread menghitung notifikasi yang belum dibaca
func (r *notificationRepositoryImpl) CountUnread(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("user_id = ? AND read_at IS NULL", userID).
		Count(&count).Error
	return count, err
}
