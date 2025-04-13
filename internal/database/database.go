package database

import (
	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := config.BuildDSN(config.CFG)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Could not connect with the database")
	}
}

func Migrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Block{},
		&models.Announcement{},
		&models.House{},
		&models.IplPayment{}, // Add IPLPayment model to MI
		&models.Post{},
		&models.PostComment{},
		&models.Resident{},
		&models.RondaGroup{},
		&models.RondaGroupMember{},
		&models.RondaContributionItem{},
		&models.Shop{},
		&models.ShopProduct{},
	)
}
