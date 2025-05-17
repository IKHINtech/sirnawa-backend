package services

import (
	"errors"

	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"gorm.io/gorm"
)

type FCMTokenService struct {
	tokenRepo *repository.FCMTokenRepository
	db        *gorm.DB
}

func NewFCMTokenService(tokenRepo *repository.FCMTokenRepository, db *gorm.DB) *FCMTokenService {
	return &FCMTokenService{tokenRepo: tokenRepo, db: db}
}

func (s *FCMTokenService) withTransaction(fn func(tx *gorm.DB) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *FCMTokenService) RegisterToken(userID, token, deviceID, deviceType, appVersion, osVersion string) error {
	// Validasi input
	if token == "" || deviceID == "" {
		return errors.New("token and deviceID must be provided")
	}

	err := s.withTransaction(func(tx *gorm.DB) error {
		err := s.tokenRepo.UpsertToken(tx, userID, token, deviceID, deviceType, appVersion, osVersion)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (s *FCMTokenService) GetUserActiveTokens(userID string) ([]string, error) {
	return s.tokenRepo.GetActiveTokens(userID)
}

func (s *FCMTokenService) RemoveToken(token string) error {
	err := s.withTransaction(func(tx *gorm.DB) error {
		return s.tokenRepo.DeactivateToken(tx, token)
	})
	return err
}

func (s *FCMTokenService) CleanupTokens() error {
	err := s.withTransaction(func(tx *gorm.DB) error {
		if err := s.tokenRepo.CleanupExpiredTokens(tx); err != nil {
			return err
		}

		return s.tokenRepo.DeactivateOldInactiveTokens(tx)
	})
	return err
}
