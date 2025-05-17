package repository

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"gorm.io/gorm"
)

type FCMTokenRepository struct {
	db *gorm.DB
}

func NewFCMTokenRepository(db *gorm.DB) *FCMTokenRepository {
	return &FCMTokenRepository{db: db}
}

// UpsertToken - Insert atau update token
func (r *FCMTokenRepository) UpsertToken(tx *gorm.DB, userID, token, deviceID, deviceType, appVersion, osVersion string) error {
	newToken := &models.UserFCMToken{
		UserID:     userID,
		Token:      token,
		DeviceID:   deviceID,
		DeviceType: deviceType,
		AppVersion: appVersion,
		OSVersion:  osVersion,
		IsActive:   true,
	}

	// Nonaktifkan token lama untuk device yang sama
	if err := tx.Model(&models.UserFCMToken{}).
		Where("user_id = ? AND device_id = ? AND token != ?", userID, deviceID, token).
		Update("is_active", false).Error; err != nil {
		return err
	}

	// Upsert token baru
	return tx.Where(models.UserFCMToken{UserID: userID, DeviceID: deviceID}).
		Assign(newToken).
		FirstOrCreate(&models.UserFCMToken{}).Error
}

// GetActiveTokens - Ambil semua token aktif user
func (r *FCMTokenRepository) GetActiveTokens(userID string) ([]string, error) {
	var tokens []string
	err := r.db.Model(&models.UserFCMToken{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Pluck("token", &tokens).Error
	return tokens, err
}

// DeactivateToken - Nonaktifkan token tertentu
func (r *FCMTokenRepository) DeactivateToken(tx *gorm.DB, token string) error {
	return tx.Model(&models.UserFCMToken{}).
		Where("token = ?", token).
		Updates(map[string]any{
			"is_active":  false,
			"updated_at": time.Now(),
		}).Error
}

// CleanupExpiredTokens - Hapus token yang sudah expired
func (r *FCMTokenRepository) CleanupExpiredTokens(tx *gorm.DB) error {
	// Hapus token yang expired lebih dari 7 hari
	return tx.Unscoped(). // Gunakan Unscoped untuk hard delete
				Where("expires_at < ?", time.Now().Add(-7*24*time.Hour)).
				Delete(&models.UserFCMToken{}).Error
}

// DeactivateOldInactiveTokens - Nonaktifkan token yang tidak aktif >3 bulan
func (r *FCMTokenRepository) DeactivateOldInactiveTokens(tx *gorm.DB) error {
	return tx.Model(&models.UserFCMToken{}).
		Where("is_active = ? AND updated_at < ?", false, time.Now().Add(-3*30*24*time.Hour)).
		Update("is_active", false).Error
}
