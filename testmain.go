package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"log"
	"fmt"
	//"github.com/gomodule/redigo/redis"
	//"github.com/gin-gonic/contrib/sessions"
	//"html/template"
	"github.com/mntky/foruka-go/testmodels"
)

//top page
func Top(c *gin.Context) {
	c.HTML(200, "top.tmpl", gin.H{
		"id":		"aa",
		"user":		"aa",
	})
}

//login page 
func Login(c *gin.Context) {
	c.HTML(200, "login.tmpl", nil)
}

//login process
func Auth(c *gin.Context) {
	c.Request.ParseForm()
	fmt.Println(c.Request.Form["mail"])
	fmt.Println(c.Request.Form["password"])
}

func main() {
	testmodels.initCache()

	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/css", "./css")

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	log.Datal(http.ListenAndServe(":8080", nil))
}
