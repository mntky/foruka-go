package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/mntky/foruka-go/models"
	"github.com/gomodule/redigo/redis"
	//"fmt"
)

type POST struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

func Signup(g *gin.Context) {
	g.HTML(200, "signup.tmpl", nil)
}

func Register(g *gin.Context){
	var data POST
	g.BindJSON(&data)

	pool := models.NewPool(1)
	redi := pool.Get() //connect redis
	p, err := redis.String(redi.Do("GET", data.Username))
	if p != "" {
		return
	}

	potato := data.Username + data.Password
	passhash, salt := models.Create(data.Password, potato)

	_, err = redi.Do("SET", data.Username, passhash)
	if err != nil {
		return
	}
	redi.Close()

	pool = models.NewPool(2)
	redi = pool.Get() //connect redis
	_, err = redi.Do("SET", salt, data.Username)
	if err != nil {
		return
	}
	redi.Close()
}
