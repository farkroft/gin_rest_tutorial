package controllers

import (
	"github.com/jinzhu/gorm"
)

// Controller struct
type Controller struct {
	db *gorm.DB
}

// New is initialize of db by reference
func New(db *gorm.DB) *Controller {
	return &Controller{
		db: db,
	}
}
