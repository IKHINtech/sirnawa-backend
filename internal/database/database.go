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

	dsn := config.BuildDSN(config.AppConfig)

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
		// &models.User{},
		// &models.UserVerification{},
		// &models.Block{},
		// &models.Announcement{},
		// &models.House{},
		// &models.IplPayment{},
		// &models.Post{},
		// &models.PostComment{},
		// &models.Resident{},
		// &models.RondaGroup{},
		// &models.RondaActivity{},
		// &models.RondaAttendance{},
		// &models.RondaGroupMember{},
		// &models.RondaConstribution{},
		// &models.RondaContributionItem{},
		// &models.Rw{},
		// &models.Rt{},
		// &models.Shop{},
		// &models.ShopProduct{},
		&models.HousingArea{},
		&models.IplBill{},
		&models.IplBillDetail{},
		&models.IplPayment{},
		&models.IplRate{},
		&models.IplRateDetail{},
		&models.Item{},
	// &models.ResidentHouse{},
	// &models.UserRT{},
	)
}
