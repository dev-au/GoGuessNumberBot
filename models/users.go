package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ChatID   int64 `gorm:"uniqueIndex"`
	Username string
}
type Debt struct {
	gorm.Model
	OwnerID   int64
	Owner     User `gorm:"foreignKey:OwnerID"`
	DebtorID  int64
	Debtor    User `gorm:"foreignKey:DebtorID"`
	DebtPrice int64
	CreatedAt time.Time
}
