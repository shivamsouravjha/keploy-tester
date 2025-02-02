package postgres

import (
	"fmt"
	"segwise/config"
	"segwise/models"

	_ "segwise/utils"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Get().DBHostReader,
		config.Get().DBUserName,
		config.Get().DBPassword,
		config.Get().DBName,
		config.Get().DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("Error connecting to PostgreSQL", zap.Error(err))
		return
	}
	err = db.AutoMigrate(&models.Trigger{}, &models.EventLog{}, &models.User{})
	if err != nil {
		zap.L().Info("Database migration failed:", zap.Any("error", err))
	}

	zap.L().Info("Connected to PostgreSQL")
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
