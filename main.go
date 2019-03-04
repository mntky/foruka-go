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

//top page
func Top(c *gin.Context) {
	c.HTML(200, "top.tmpl", nil)
}

//login page 
func Login(c *gin.Context) {
	c.HTML(200, "login2.tmpl", nil)
}

//login process
func Auth(c *gin.Context) {
	var authdata Authdata
	c.BindJSON(&authdata)
	fmt.Printf("ok %s", authdata)
	testmodels.Signin(c, authdata.Username, authdata.Password)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/css", "templates/css")

	r.GET("/",  Top)		//Top page
	r.GET("/login", Login)		//login page
	r.POST("/auth", Auth)		//login process

	r.Run(":8080")
}
