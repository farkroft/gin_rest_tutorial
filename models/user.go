package models

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	// ID        uuid.UUID `json:"id" gorm:"column:id;type:uuid;primary_key;"`
	FirstName string `json:"first_name" gorm:"column:first_name;type:varchar(100)"`
	LastName  string `json:"last_name" gorm:"column:last_name;type:varchar(100)"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"UpdatedAt"`
	// DeletedAt time.Time `json:"DeletedAt"`
	Products []*Product `json:"Product"`
	// CustomerID uint        `json:CustomerID`
	Customers []*Customer `gorm:"many2many:user_customers"`
}

// BeforeCreate func
// func (p *User) BeforeCreate(scope *gorm.Scope) error {
// 	uuid, err := uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uuid)
// }
