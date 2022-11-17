package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var DB *gorm.DB
var err error

type Todo struct {
	gorm.Model
	Name string `json:"name"`
}

func dbConnect() *gorm.DB {

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := user + ":" + password + "@tcp(mysql:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func main() {
	engine := gin.Default()
	db := dbConnect()
	db.AutoMigrate(&Todo{})

	// CREATE
	engine.POST("/todo", func(c *gin.Context) {
		todo := Todo{}
		c.BindJSON(&todo)
		if err := db.Create(&todo).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, todo)
	})

	// READ
	engine.GET("/todo/:id", func(c *gin.Context) {
		todo := Todo{}
		id := c.Param("id")
		if err := db.First(&todo, id).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, todo)
	})

	// UPDATE
	engine.PUT("/todo/:id", func(c *gin.Context) {
		todo := Todo{}
		id := c.Param("id")
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			fmt.Println(err)
		}
		c.BindJSON(&todo)
		db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	})

	// DELETE
	engine.DELETE("/todo/:id", func(c *gin.Context) {
		todo := Todo{}
		id := c.Param("id")
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			fmt.Println(err)
		}
		db.Delete(&todo)
		c.JSON(http.StatusOK, todo)
	})

	engine.Run(":8080")
}
