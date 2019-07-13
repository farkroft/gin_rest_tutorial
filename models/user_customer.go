package models

import (
	"time"
)

// UserCustomer struct
type UserCustomer struct {
	ID         uint      `gorm:"primary_key"`
	CustomerID uint      `sql:"type:int REFERENCES Customers(id)"`
	Customer   Customer  `gorm:"foreignkey:CustomerID"`
	UserID     uint      `sql:"type:int REFERENCES Users(id)"`
	User       User      `gorm:"foreignkey:UserID"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
	DeletedAt  time.Time `json:"DeletedAt"`
}
