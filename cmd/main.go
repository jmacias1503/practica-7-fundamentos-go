package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int 
	Name string 
	Email string 
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to DB")
	}
	db.AutoMigrate(&User{})
	router := gin.Default()
	var users []User
	fmt.Println("Running app")
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "My first Go website with gin",
			"total_users": len(users),
			"users": users,
		})
	})
	router.GET("/api/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
	})
	router.POST("/api/users", func(c *gin.Context) {
		var user User
		if c.BindJSON(&user) == nil {
			db.Create(&user)
			c.JSON(201, user)
		} else {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
		}
	})
	router.DELETE("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		err := db.First(&user, id).Error
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid id",
			})
		}
		db.Delete(&user)
		c.JSON(201, gin.H{"message": "User deleted"})
	})

	router.PUT("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
			var user User
			err := db.First(&user, id).Error
			if err != nil {
				c.JSON(404, gin.H{
					"error": "User not found",
				})
			}
			err = c.BindJSON(&user)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "Invalid payload",
				})
				return
			}
			db.Save(&user)
		c.JSON(201, gin.H{})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
