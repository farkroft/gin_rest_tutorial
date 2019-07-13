package main

import (
	"fmt"
	"log"
	"os"

	controllers "github.com/farkroft/gin_rest_tutorial/controllers"
	models "github.com/farkroft/gin_rest_tutorial/models"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}
	dialect := os.Getenv("DB_DIALECT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=rest_gin_tutorial password=postgres sslmode=disable")
	if dialect != "sqlite3" {
		databaseURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, username, dbName, password)
		db, err = gorm.Open(dialect, databaseURL)
	}

	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection established ...")
	// drop db when run
	// db.Debug().DropTableIfExists(&UserCustomer{})
	// db.Debug().DropTableIfExists(&Product{})
	// db.Debug().DropTableIfExists(&User{})
	// db.Debug().DropTableIfExists(&Customer{})

	db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Customer{}, &models.UserCustomer{})
	db.Model(&models.Product{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")
	db.Model(&models.UserCustomer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")
	db.Model(&models.UserCustomer{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "CASCADE")
	defer db.Close()

	r := gin.Default()
	ct := controllers.New(db)

	// Users API
	r.GET("/users", ct.GetAllUsers)
	r.GET("/user/:id", ct.GetUser)
	r.POST("/user", ct.CreateUser)
	r.PUT("/user/:id", ct.UpdateUser)
	r.DELETE("/user/:id", ct.DeleteUser)

	// Products API
	r.GET("/products", ct.GetAllProducts)
	r.GET("/product/:id", ct.GetProduct)
	r.POST("/product", ct.CreateProduct)
	r.PUT("/product/:id", ct.UpdateProduct)
	r.DELETE("/product/:id", ct.DeleteProduct)

	// Customers API
	r.GET("/customers", ct.GetAllCustomers)
	r.GET("/customer/:id", ct.GetCustomer)
	r.POST("/customer", ct.CreateCustomer)
	r.PUT("/customer/:id", ct.UpdateCustomer)
	r.DELETE("/customer/:id", ct.DeleteCustomer)

	// r.POST("/customer/user/:id", CreateUserCustomer)

	r.Run(":8080")
}
