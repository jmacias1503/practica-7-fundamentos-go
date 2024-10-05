package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  fmt.Println("Running app")
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
