package db

import (
	models2 "ecosnap/internal/app/models"
	"ecosnap/internal/app/models/seeds"
	"ecosnap/pkg/logger"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models2.User{},
		&models2.TrashCategories{},
		&models2.Trash{},
	)
	if err != nil {
		logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	if err = seeds.SeedTrashCategories(dbConn); err != nil {
		logger.Error.Printf("[db.Migrate] Error migrating trash categories: %v", err)

		return err
	}

	return nil
}
