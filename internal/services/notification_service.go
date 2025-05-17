package services

import (
	"context"
	"errors"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
	repositories "github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, input models.Notification) (*models.Notification, error)
	GetUserNotifications(pagination utils.Pagination, userID, rtID, houseID string) (*utils.Pagination, []models.Notification, error)
	MarkNotificationAsRead(ctx context.Context, id, userID string) error
	MarkAllNotificationsAsRead(ctx context.Context, userID string) error
	DeleteNotification(ctx context.Context, id, userID string) error
	GetUnreadCount(ctx context.Context, userID string) (int64, error)
}

type notificationServiceImpl struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationServiceImpl{repo: repo}
}

// CreateNotification membuat notifikasi baru
func (s *notificationServiceImpl) CreateNotification(ctx context.Context, input models.Notification) (*models.Notification, error) {
	if input.UserID == "" {
		return nil, errors.New("user ID is required")
	}

	if input.Title == "" || input.Body == "" {
		return nil, errors.New("title and body are required")
	}

	err := s.repo.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}

// GetUserNotifications mendapatkan notifikasi user
func (s *notificationServiceImpl) GetUserNotifications(pagination utils.Pagination, userID, rtID, houseID string) (*utils.Pagination, []models.Notification, error) {
	return s.repo.GetByUserID(pagination, userID, rtID, houseID)
}

// MarkNotificationAsRead menandai notifikasi sebagai dibaca
func (s *notificationServiceImpl) MarkNotificationAsRead(ctx context.Context, id, userID string) error {
	// Verifikasi notifikasi milik user
	notification, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if notification.UserID != userID {
		return errors.New("unauthorized: notification does not belong to user")
	}

	return s.repo.MarkAsRead(ctx, id)
}

// MarkAllNotificationsAsRead menandai semua notifikasi user sebagai dibaca
func (s *notificationServiceImpl) MarkAllNotificationsAsRead(ctx context.Context, userID string) error {
	return s.repo.MarkAllAsRead(ctx, userID)
}

// DeleteNotification menghapus notifikasi
func (s *notificationServiceImpl) DeleteNotification(ctx context.Context, id, userID string) error {
	// Verifikasi notifikasi milik user
	notification, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if notification.UserID != userID {
		return errors.New("unauthorized: notification does not belong to user")
	}

	return s.repo.Delete(ctx, id)
}

// GetUnreadCount mendapatkan jumlah notifikasi yang belum dibaca
func (s *notificationServiceImpl) GetUnreadCount(ctx context.Context, userID string) (int64, error) {
	return s.repo.CountUnread(ctx, userID)
}
