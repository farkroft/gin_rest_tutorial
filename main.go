package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

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

// Product struct
type Product struct {
	gorm.Model
	Brand  string `json:"Brand" gorm:"type:varchar(100)"`
	Name   string `json:"Name" gorm:"type:varchar(100)"`
	UserID uint   `json:"UserID"`
	User   *User
}

// Customer Struct
type Customer struct {
	gorm.Model
	FirstName string  `json:"first_name" gorm:"column:first_name;type:varchar(100)"`
	LastName  string  `json:"last_name" gorm:"column:last_name;type:varchar(100)"`
	Users     []*User `gorm:"many2many:user_customers"`
	UserID    uint    `json:"UserID"`
}

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

func main() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=rest_gin_tutorial password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connection established ...")
	// db.Debug().DropTableIfExists("users_customers")
	// db.Debug().DropTableIfExists(&UserCustomer{})
	// db.Debug().DropTableIfExists(&Product{})
	// db.Debug().DropTableIfExists(&User{})
	// db.Debug().DropTableIfExists(&Customer{})

	db.Debug().AutoMigrate(&User{}, &Product{}, &Customer{}, &UserCustomer{})
	db.Model(&Product{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")
	db.Model(&UserCustomer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")
	db.Model(&UserCustomer{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "CASCADE")
	defer db.Close()

	r := gin.Default()

	// Users API
	r.GET("/users", GetAllUsers)
	r.GET("/user/:id", GetUser)
	r.POST("/user", CreateUser)
	r.PUT("/user/:id", UpdateUser)
	r.DELETE("/user/:id", DeleteUser)

	// Products API
	r.GET("/products", GetAllProducts)
	r.GET("/product/:id", GetProduct)
	r.POST("/product", CreateProduct)
	r.PUT("/product/:id", UpdateProduct)
	r.DELETE("/product/:id", DeleteProduct)

	// Customers API
	r.GET("/customers", GetAllCustomers)
	r.GET("/customer/:id", GetCustomer)
	r.POST("/customer", CreateCustomer)
	r.PUT("/customer/:id", UpdateCustomer)
	r.DELETE("/customer/:id", DeleteCustomer)

	// r.POST("/customer/user/:id", CreateUserCustomer)

	r.Run(":8080")
}

// BeforeCreate func
// func (p *User) BeforeCreate(scope *gorm.Scope) error {
// 	uuid, err := uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uuid)
// }

// CreateUser to create new user
func CreateUser(c *gin.Context) {
	var users User
	c.BindJSON(&users)

	db.Create(&users)
	c.JSON(http.StatusOK, gin.H{"result": users, "error": nil})
	// c.JSON(200, users)
}

// GetAllUsers func to run when root page accessed
func GetAllUsers(c *gin.Context) {
	users := []User{}

	User := db.Debug().Preload("Products").Preload("Customers").Find(&users)
	// fmt.Println(User)
	if err := db.Debug().Preload("Products").Preload("Customers").Find(&users).Error; err != nil {
		// if err := db.Find(&User).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": User.Value, "error": nil, "status": http.StatusOK})
	}
}

// GetUser func to get a users details
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)

	users := &User{}

	result := db.Where("id = ?", id).First(&users)
	// result := db.Where("first_name = ?", "Jhonny").First(users)
	if result.Error != nil {
		c.AbortWithStatus(404)
		fmt.Println(result.Error.Error())
	} else {
		c.JSON(200, &users)
	}
}

// UpdateUser func to update a users detail
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var users User

	if err := db.Where("id = ?", id).First(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&users)
	db.Save(&users)
	c.JSON(200, users)
}

// DeleteUser func to delete a users
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var users User

	if err := db.Where("id = ?", id).First(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	d := db.Where("id = ?", id).Delete(&users)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// GetAllProducts func to run when root page accessed
func GetAllProducts(c *gin.Context) {
	var products []Product
	Product := db.Find(&products)
	// Product := db.Debug().Preload("User").Find(&products) // to load relation
	if err := Product.Error; err != nil {
		// if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, Product)
	}
}

// CreateProduct to create new users
func CreateProduct(c *gin.Context) {
	var products Product
	c.BindJSON(&products)

	db.Create(&products)
	c.JSON(200, products)
}

// GetProduct func to get a users details
func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var products Product
	if err := db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

// UpdateProduct func to update a users detail
func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var products Product

	if err := db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&products)
	db.Save(&products)
	c.JSON(200, products)
}

// DeleteProduct func to delete a users
func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var products Product

	if err := db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	d := db.Where("id = ?", id).Delete(&products)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// GetAllCustomers func to run when root page accessed
func GetAllCustomers(c *gin.Context) {
	var customers []Customer

	Customers := db.Debug().Preload("Users").Find(&customers)
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
func GetCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customers Customer
	if err := db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customers)
	}
}

// CreateCustomer to create new users
func CreateCustomer(c *gin.Context) {
	var customers Customer

	c.BindJSON(&customers)
	db.Create(&customers)

	fmt.Println("CustomerID ", customers.ID)
	if customers.UserID > 0 {
		db.Create(&UserCustomer{
			CustomerID: customers.ID,
			UserID:     customers.UserID,
		})
	}

	c.JSON(200, customers)
}

// UpdateCustomer func to update a users detail
func UpdateCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customers Customer

	if err := db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&customers)
	db.Save(&customers)
	if customers.UserID > 0 {
		db.Create(&UserCustomer{
			CustomerID: customers.ID,
			UserID:     customers.UserID,
		})
	}
	c.JSON(200, customers)
}

// DeleteCustomer func to delete a users
func DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customers Customer

	if err := db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	d := db.Where("id = ?", id).Delete(&customers)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
