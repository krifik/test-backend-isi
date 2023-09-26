package entity

type Account struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	UserID        int    `gorm:"not null;default:0"`
	AccountNumber string `gorm:"not null"`
	Balance       int64  `gorm:"not null;default:0"`
}
