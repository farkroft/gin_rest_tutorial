package models

import "github.com/jinzhu/gorm"

// Product struct
type Product struct {
	gorm.Model
	Brand  string `json:"Brand" gorm:"type:varchar(100)"`
	Name   string `json:"Name" gorm:"type:varchar(100)"`
	UserID uint   `json:"UserID"`
	User   *User
}
