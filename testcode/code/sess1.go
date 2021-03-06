package testmodels

import (
	//"log"
	"net/http"
	"github.com/gomodule/redigo/redis"
	"fmt"
	//"encoding/json"
	"github.com/satori/go.uuid"
	"time"
	"github.com/gin-gonic/gin"
)

var cache redis.Conn

var user = map[string]string {
	"user1":	"123qwEcc",
	"uesr2":	"123qwEcc@",
}


func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}

func Signin(g *gin.Context, uname, pass string) {
	InitCache()
	fmt.Println("ok1")
/*
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Println(r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
*/
	getpass, ok := user[uname]

	if !ok || getpass != pass {
		return
	}

	fmt.Println("ok2")

	u, err := uuid.NewV4()
	sessionToken := u.String()

	_, err = cache.Do("SETEX", sessionToken, "120", uname)
	if err != nil {
		return
	}

	fmt.Println("ok3")

	http.SetCookie(g.Writer, &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,
		Expires:	time.Now().Add(120 * time.Second),
	})
	g.Redirect(http.StatusMovedPermanently,"http://192.168.11.100/welcome")
}

func Welcome(g *gin.Context) {
	c, err := g.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return 
		}
		return 
	}
	sessionToken := c

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		return 
	}
	if response == nil {
		return 
	}

	g.HTML(200, "success.tmpl", nil)
}


