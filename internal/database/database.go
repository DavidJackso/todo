package database

import (
	"fmt"
	"os"

	"github.com/DavidJackso/TodoApi/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(cfg *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Address,
		cfg.User,
		os.Getenv("POSTGRES_PASSWORD"),
		cfg.Name,
		cfg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("failed connect to db: %v", err)
		err := fmt.Errorf("failed connect tod db")
		return nil, err
	}

	return db, nil
}
