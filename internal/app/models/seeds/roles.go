package seeds

import (
	"ecosnap/internal/app/models"
	"ecosnap/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func SeedTrashCategories(db *gorm.DB) error {
	trs := []models.TrashCategories{
		{ID: 1, Name: "Пластик"},
		{ID: 2, Name: "Макулатура / Бумага"},
		{ID: 3, Name: "Стекло"},
		{ID: 4, Name: "Металл"},
		{ID: 5, Name: "Тетрапак"},
	}

	for _, tr := range trs {
		var existingCategorie models.TrashCategories
		if err := db.First(&existingCategorie, "name = ?", tr.Name).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&tr)
			} else {
				logger.Error.Printf("[seeds.SeedTrashCategories] Error seeding ct: %v", err)

				return err
			}
		}
	}

	return nil
}
