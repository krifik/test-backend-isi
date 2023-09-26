package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID                     int    `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	Name                   string `gorm:"column:name"`
	NationalIdentityNumber string `gorm:"column:national_identity_number"`
	PhoneNumber            string `gorm:"column:phone_number"`
}
