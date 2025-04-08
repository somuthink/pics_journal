package db

import (
	"fmt"

	"github.com/somuthink/pics_journal/core/internal/config"
	"github.com/somuthink/pics_journal/core/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/charmbracelet/log"
)

var DB *gorm.DB

func DSN(dbName string) string {
	cfg := config.Cfg
	base := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD)
	if dbName != "" {
		return fmt.Sprintf("%s dbname=%s", base, dbName)
	}
	return fmt.Sprintf("%s dbname=postgres", base)
}

func OpenOrCreateDialcector(dbName string) (gorm.Dialector, error) {
	dsn := DSN("")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	cnt := 0
	db.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&cnt)
	if cnt == 0 {
		sql := fmt.Sprintf("CREATE DATABASE %s", dbName)
		log.Info("DB init", "creating new db", "name", dbName)
		result := db.Exec(sql)
		if result.Error != nil {
			return nil, err
		}
	}

	return postgres.Open(DSN(dbName)), nil
}

func Initialize() error {
	dialector, err := OpenOrCreateDialcector(config.Cfg.DB_NAME)
	if err != nil {
		return err
	}
	DB, err = gorm.Open(dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return err
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm")

	if err := DB.AutoMigrate(
		&models.User{}, &models.Input{}); err != nil {
		return err
	}

	return nil
}
