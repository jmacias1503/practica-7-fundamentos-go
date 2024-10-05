package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	router := gin.Default()
	users := []User {
		User{
			Id: "snoatehueoa",
			Name: "Alfredo",
			Email: "alfredo@mail.com",
		},
	}
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
	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
