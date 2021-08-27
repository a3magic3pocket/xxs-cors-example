package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:8888"},
			AllowMethods:     []string{"POST"},
			AllowHeaders:     []string{"Origin", "Cookie", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	router.POST("trap", func(c *gin.Context) {
		fmt.Println("Cookie", c.Request.Header["Cookie"])
		fmt.Println("Authorization", c.Request.Header["Authorization"])
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
	})

	router.Run(":8889")
}
