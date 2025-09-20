package models

type TrashCategories struct {
	ID uint `gorm:"primarykey"`

	Name string `gorm:"type:varchar(50);not null"`
}
