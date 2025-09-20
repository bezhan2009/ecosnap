package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(255);unique;not null"`
	Email    string `gorm:"type:varchar(255);unique;"`
	Phone    string `gorm:"type:varchar(9);unique;"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`

	FullName string `json:"full_name" gorm:"type:varchar(255);"`
}
