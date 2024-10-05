package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
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
		c.HTML(200, "index.html", gin.H{})
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
	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
