package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mntky/foruka-go/testmodels"

	"fmt"
)

func main() {
	r := gin.Default()

	//Load files
	r.LoadHTMLGlob("view/*")
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/css", "./css")

	//Route to Controller
	r.GET("/", controllers.Home)			//Top page
	r.GET("/login",	controllers.Login)		//Login page
	r.POST("/auth", controllers.Auth)		//Login process
	r.GET("/welcome", controlers.Welcome)	//Sucess Login

	r.Run(":8080")
}
