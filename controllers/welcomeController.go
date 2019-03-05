package controllers

import(
	"github.com/gin-gonic/gin"
)

func Welcome(g *gin.Context) {
	g.HTML(200, "welcome.tmpl", nil)
}
