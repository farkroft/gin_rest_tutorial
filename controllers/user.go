package controllers

import (
	"fmt"
	"net/http"

	"github.com/farkroft/gin_rest_tutorial/models"
	"github.com/gin-gonic/gin"
)

// CreateUser to create new user
func (ct *Controller) CreateUser(c *gin.Context) {
	var users models.User
	c.BindJSON(&users)

	ct.db.Create(&users)
	c.JSON(http.StatusOK, gin.H{"result": users, "error": nil})
	// c.JSON(200, users)
}

// GetAllUsers func to run when root page accessed
func (ct *Controller) GetAllUsers(c *gin.Context) {
	users := []models.User{}

	User := ct.db.Debug().Preload("Products").Preload("Customers").Find(&users)
	// fmt.Println(User)
	if err := ct.db.Debug().Preload("Products").Preload("Customers").Find(&users).Error; err != nil {
		// if err := ct.db.Find(&User).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": User.Value, "error": nil, "status": http.StatusOK})
	}
}

// GetUser func to get a users details
func (ct *Controller) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)

	users := &models.User{}

	result := ct.db.Where("id = ?", id).First(&users)
	// result := ct.db.Where("first_name = ?", "Jhonny").First(users)
	if result.Error != nil {
		c.AbortWithStatus(404)
		fmt.Println(result.Error.Error())
	} else {
		c.JSON(200, &users)
	}
}

// UpdateUser func to update a users detail
func (ct *Controller) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var users models.User

	if err := ct.db.Where("id = ?", id).First(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&users)
	ct.db.Save(&users)
	c.JSON(200, users)
}

// DeleteUser func to delete a users
func (ct *Controller) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var users models.User

	if err := ct.db.Where("id = ?", id).First(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	d := ct.db.Where("id = ?", id).Delete(&users)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
