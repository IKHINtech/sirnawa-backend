package helpers

import (
	"errors"

	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"gorm.io/gorm"
)

func GetUserByEmail(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email: &e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
