package database

import (
	"fmt"
	"os"

	"github.com/DavidJackso/TodoApi/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBConfig.Address,
		cfg.DBConfig.User,
		os.Getenv("POSTGRES_PASSWORD"),
		cfg.DBConfig.Name,
		cfg.DBConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to connect database")
	}
	return db
}
