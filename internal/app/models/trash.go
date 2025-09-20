package models

import "gorm.io/gorm"

type Trash struct {
	gorm.Model

	UserID uint `json:"user_id"`
	User   User `json:"-" gorm:"foreignKey:UserID"`

	TrashCategoriesID uint            `json:"trash_categories_id"`
	TrashCategories   TrashCategories `json:"trash_categories" gorm:"foreignKey:TrashCategoriesID"`
}
