package repository

import (
	"ecosnap/internal/app/models"
	"ecosnap/pkg/db"
	"ecosnap/pkg/logger"
)

func GetAllTrashes(month, year int) (trashes []models.Trash, err error) {
	if err = db.GetDBConn().Model(&models.Trash{}).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Find(&trashes).Error; err != nil {
		logger.Error.Printf("[repository.GetAllTrashes] Error getting trashes: %s", err.Error())

		return nil, TranslateGormError(err)
	}

	return trashes, nil
}

func GetTrashByID(trashID int) (trashes models.Trash, err error) {
	if err = db.GetDBConn().Model(&models.Trash{}).
		Where("id = ?", trashID).
		First(&trashes).Error; err != nil {
		logger.Error.Printf("[repository.GetAllTrashes] Error getting trashes: %s", err.Error())

		return models.Trash{}, TranslateGormError(err)
	}

	return trashes, nil
}

func CreateTrash(trash models.Trash) (err error) {
	if err = db.GetDBConn().Model(&models.Trash{}).Create(&trash).Error; err != nil {
		logger.Error.Printf("[repository.CreateTrash] Error creating trash: %s", err.Error())

		return TranslateGormError(err)
	}

	return nil
}

func DeleteTrash(trashID int) (err error) {
	if err = db.GetDBConn().Model(&models.Trash{}).Delete(&models.Trash{}, trashID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteTrash] Error deleting trash: %s", err.Error())

		return TranslateGormError(err)
	}

	return nil
}
