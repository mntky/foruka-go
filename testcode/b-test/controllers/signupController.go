package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/mntky/foruka-go/models"
	"github.com/gomodule/redigo/redis"
	"fmt"
	//"net/http"
	"encoding/json"
)

type POST struct {
	Username	string
	Password	string
}

func Signup(g *gin.Context) {
	g.HTML(200, "signup.tmpl", nil)
}

func Register(g *gin.Context){
	defer g.Redirect(301, "/login")
	defer g.Abort()

	g.Request.ParseForm()
	fmt.Println(g.Request.Form["username"])
	fmt.Println(g.Request.Form["passowrd"])


	j := []byte(`{"Username":g.Param("username"),"Password":g.Param("password")}`)
	var pos POST
	json.Unmarshal(j, &pos)

	fmt.Println(&pos)
	data := pos

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
