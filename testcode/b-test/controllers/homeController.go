package controllers

import (
	"github.com/gin-gonic/gin"
)

func Top(c *gin.Context) {
	c.HTML(200, "top.tmpl", nil)
}
