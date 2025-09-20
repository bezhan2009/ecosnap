package service

import (
	"ecosnap/internal/app/models"
	"ecosnap/internal/repository"
)

func GetAllTrashes(month, year int) (trashes []models.Trash, err error) {
	return repository.GetAllTrashes(month, year)
}

func GetTrashByID(trashID int) (trash models.Trash, err error) {
	return repository.GetTrashByID(trashID)
}

func CreateTrash(trash models.Trash) (err error) {
	return repository.CreateTrash(trash)
}

func DeleteTrash(trashID int) (err error) {
	return repository.DeleteTrash(trashID)
}
