package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	//"log"
	"fmt"
	//"github.com/gin-gonic/contrib/sessions"
	//"html/template"
	"github.com/mntky/foruka-go/testmodels"
)

type Authdata struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/css", "templates/css")

	r.GET("/",  controllers.Top)					//Top page
	r.GET("/login", controllers.Login)				//login page
	r.POST("/auth", controllers.Auth)				//login process
	r.GET("/welcome", controllers.Welcome)	//login process

	r.Run(":8080")
}
