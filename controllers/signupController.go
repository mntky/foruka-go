package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/mntky/foruka-go/models"
	"github.com/gomodule/redigo/redis"
	//"fmt"
	"strings"
	//"net/http"
)

func Signup(g *gin.Context) {
	g.HTML(200, "signup.tmpl", nil)
}

func Register(g *gin.Context){
	var uname, passwd string
	defer g.Redirect(301, "/login")
	defer g.Abort()

	g.Request.ParseForm()
	uname = strings.Join(g.Request.PostForm["username"]," ") 
	passwd = strings.Join(g.Request.PostForm["password"]," ") 

	pool := models.NewPool(1)
	redi := pool.Get() //connect redis
	p, err := redis.String(redi.Do("GET", uname))
	if p != "" {
		return
	}

	potato := uname + passwd
	passhash, salt := models.Create(passwd, potato)

	_, err = redi.Do("SET", uname, passhash)
	if err != nil {
		return
	}
	redi.Close()

	pool = models.NewPool(2)
	redi = pool.Get() //connect redis
	_, err = redi.Do("SET", salt, uname)
	if err != nil {
		return
	}
	redi.Close()
}
