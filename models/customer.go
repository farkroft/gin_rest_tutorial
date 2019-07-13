package models

import "github.com/jinzhu/gorm"

// Customer Struct
type Customer struct {
	gorm.Model
	FirstName string  `json:"first_name" gorm:"column:first_name;type:varchar(100)"`
	LastName  string  `json:"last_name" gorm:"column:last_name;type:varchar(100)"`
	Users     []*User `gorm:"many2many:user_customers"`
	UserID    uint    `json:"UserID"`
}
