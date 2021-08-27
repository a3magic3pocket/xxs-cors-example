package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("view/*")
	router.GET("victim/community/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "victim-community.html", gin.H{})
	})
	router.GET("victim/community/kakao", func(c *gin.Context) {
		c.HTML(http.StatusOK, "victim-community-kakao.html", gin.H{})
	})
	router.Run(":8888")
}
