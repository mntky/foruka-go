package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
	//"github.com/gin-gonic/contrib/sessions"
	//"html/template"
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
func Auth(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == "GET" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
	fmt.Println(body)
	}
}



func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/css", "./css")

	r.GET("/",  Top)			//Top page
	r.GET("/login", Login)		//login page
	r.POST("/auth", Auth)		//login process

	r.Run(":8080")
}
