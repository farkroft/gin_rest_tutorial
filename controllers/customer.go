package controllers

import (
	"fmt"
	"net/http"

	"github.com/farkroft/gin_rest_tutorial/models"
	"github.com/gin-gonic/gin"
)

// GetAllCustomers func to run when root page accessed
func (ct *Controller) GetAllCustomers(c *gin.Context) {
	var customers []models.Customer

	Customers := ct.db.Debug().Preload("Users").Find(&customers)
	// fmt.Println(User)
	if err := Customers.Error; err != nil {
		// if err := db.Find(&User).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": Customers.Value, "error": nil, "status": http.StatusOK})
	}
}

// GetCustomer func to get a users details
func (ct *Controller) GetCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customers models.Customer
	if err := ct.db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customers)
	}
}

// CreateCustomer to create new users
func (ct *Controller) CreateCustomer(c *gin.Context) {
	var customers models.Customer

	c.BindJSON(&customers)
	ct.db.Create(&customers)

	fmt.Println("CustomerID ", customers.ID)
	if customers.UserID > 0 {
		ct.db.Create(&models.UserCustomer{
			CustomerID: customers.ID,
			UserID:     customers.UserID,
		})
	}

	c.JSON(200, customers)
}

// UpdateCustomer func to update a users detail
func (ct *Controller) UpdateCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customers models.Customer

	if err := ct.db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customers)
	ct.db.Save(&customers)
	if customers.UserID > 0 {
		ct.db.Create(&models.UserCustomer{
			CustomerID: customers.ID,
			UserID:     customers.UserID,
		})
	}
	c.JSON(200, customers)
}

// DeleteCustomer func to delete a users
func (ct *Controller) DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customers models.Customer

	if err := ct.db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	d := ct.db.Where("id = ?", id).Delete(&customers)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
