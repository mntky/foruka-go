package controllers

import (
	"github.com/mntky/foruka-go/models"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"time"
	"fmt"
)

type POSTdata struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
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

	//check passwordd with redis
	pool := models.NewPool(1)
	redi := pool.Get()
	//fmt.Println(data.Username)
	p, err := redis.String(redi.Do("GET", data.Username))
	//fmt.Println(p)
	if err != nil {
		return
	}

	potato := data.Username + data.Password
	passhash, salt := models.Create(data.Password, potato)
	p, err = redis.String(redi.Do("GET", data.Username))
	if passhash != p {
		fmt.Println(salt)
		return
	}
	redi.Close()

	//Create token of cookie, and it in sessionToken
	u, err := uuid.NewV4()
	sessionToken := u.String()

	//Set SessionToken to Redis 
	pool = models.NewPool(0)
	redi = pool.Get()
	_, err = redi.Do("SETEX", sessionToken, "1800", data.Username)
	if err != nil {
		return
	}
	redi.Close()

	//Set Cookie to User
	http.SetCookie(g.Writer, &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,
		Expires:	time.Now().Add(120 * time.Second),
	})
}
