package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/mntky/foruka-go/models"
	"net/http"
)

func Welcome(g *gin.Context) {
	cache := models.InitCache()

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

	g.HTML(200, "welcome.tmpl", nil)
}
