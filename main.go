package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*") 
	r.Static("/css", "./css")

	//home/
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "foruka-go",
		})
	})

	//user
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
		})
	})

	//login
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
		})
	})

	//lxd-controller
	r.GET("/lxd", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
		})
	})


	r.Run(":8080")
}
