package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")

	r.Static("/css", "./public")

	//http://192.168.11.100:8080/
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
		gin.H{
			"title": "foruka-go",
		})
	})
		/*
		c.JSON(200, gin.H{
			"message": "pong",
		})
		*/

	r.Run(":8080")
}
