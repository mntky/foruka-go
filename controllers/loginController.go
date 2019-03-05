package controllers

import (
	"github.com/mntky/foruka-go/models"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/mntky/foruka-go/models"
	"net/http"
	"time"
	//"fmt"
)

type POSTdata struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

var user = map[string]string {
	"user1":	"123qwEcc",
	"user2":	"123qwEcc@",
}

//Default Login Page
func Login(g *gin.Context) {
	g.HTML(200, "login2.tmpl", nil)
}

//receive Post data from login2.tmpl
//processing Posted JSON data(username, password) and Send to SigninFunc
func Auth(g *gin.Context) {
	var data POSTdata
	g.BindJSON(&data)

	cache := models.InitCache()

	//check password
	getpass, check := user[data.Username]
	if !check || getpass != data.Password {
		return
	}

	//Create token of cookie, and it in sessionToken
	u, err := uuid.NewV4()
	sessionToken := u.String

	//Set SessionToken to Redis 
	_, err = cache.Do("SETEX", sessionToken, "120", data.Username)
	if err != nil {
		return
	}

	//Set Cookie to User
	http.SetCookie(g.Writer, &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,
		Expires:	time.Now().Add(120 * time.Second),
	})

	return
}
