package controllers

import (
	"fmt"

	"github.com/farkroft/gin_rest_tutorial/models"
	"github.com/gin-gonic/gin"
)

// GetAllProducts func to run when root page accessed
func (ct *Controller) GetAllProducts(c *gin.Context) {
	var products []models.Product
	Product := ct.db.Find(&products)
	// Product := ct.db.Debug().Preload("User").Find(&products) // to load relation
	if err := Product.Error; err != nil {
		// if err := ct.db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, Product)
	}
}

// CreateProduct to create new users
func (ct *Controller) CreateProduct(c *gin.Context) {
	var products models.Product
	c.BindJSON(&products)

	ct.db.Create(&products)
	c.JSON(200, products)
}

// GetProduct func to get a users details
func (ct *Controller) GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var products models.Product
	if err := ct.db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

// UpdateProduct func to update a users detail
func (ct *Controller) UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var products models.Product

	if err := ct.db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&products)
	ct.db.Save(&products)
	c.JSON(200, products)
}

// DeleteProduct func to delete a users
func (ct *Controller) DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var products models.Product

	if err := ct.db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	d := ct.db.Where("id = ?", id).Delete(&products)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
