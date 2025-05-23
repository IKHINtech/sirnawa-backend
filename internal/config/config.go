package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	DBSSLMode        string `mapstructure:"DB_SSLMODE"`
	PORT             string `mapstructure:"PORT"`
	JWT_SECRET       string `mapstructure:"JWT_SECRET"`
	DEFAULT_PASSWORD string `mapstructure:"DEFAULT_PASSWORD"`
	EMAIL_HOST       string `mapstructure:"EMAIL_HOST"`
	EMAIL_PORT       string `mapstructure:"EMAIL_PORT"`
	EMAIL_USERNAME   string `mapstructure:"EMAIL_USERNAME"`
	EMAIL_PASSWORD   string `mapstructure:"EMAIL_PASSWORD"`
	EMAIL_FROM_NAME  string `mapstructure:"EMAIL_FROM_NAME"`
	DRIVE_FOLDER     string `mapstructure:"DRIVE_FOLDER"`
}

var AppConfig Config

func LoadConfig() (config Config, err error) {
	// Cari lokasi .env (berdasarkan working directory)
	envPath := filepath.Join(".", ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("File .env tidak ditemukan di: %s", envPath)
	}

	viper.SetConfigFile(envPath)
	viper.AutomaticEnv() // Juga baca dari environment variables sistem

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	AppConfig = config
	return config, err
}

func BuildDSN(cfg Config) string {
	return "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + cfg.DBPort +
		" sslmode=" + cfg.DBSSLMode
}
