package controllers

import(
	"github.com/gin-gonic/gin"
)

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
	gi response == nil {
		return
	}

	g.HTML(200, "welcome.tmpl", nil)
}
