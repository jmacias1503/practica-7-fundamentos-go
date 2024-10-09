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
	indexUser := 1
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
		c.JSON(200, users)
	})
	router.POST("/api/users", func(c *gin.Context) {
		var user User
		if c.BindJSON(&user) == nil {
			user.Id = indexUser
			users = append(users, user)
			indexUser++
		} else {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
		}
	})
	router.DELETE("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid id",
			})
		}
		fmt.Println("Id a borrar: ", id)
		for i, user := range users {
			if user.Id ==  idParsed {
				users = append(users[:i], users[i+1:]...)
				c.JSON(200, gin.H{
					"message": "User deleted",
				})
				return
			}
		}
		c.JSON(201, gin.H{})
	})

	router.PUT("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid id",
			})
			return
		}
			var user User
			err = c.BindJSON(&user)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "Invalid payload",
				})
				return
			}
		fmt.Println("Id a actualizar: ", id)
		for i, u := range users {
			if u.Id ==  idParsed {
				users[i] = user
				users[i].Id = idParsed
				c.JSON(200, users[i])
				return
			}
		}
		c.JSON(201, gin.H{})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
