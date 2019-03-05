package controllers

import (
	"github.com/mntky/foruka-go/models"
)

func Top(c *gin.Context) {
	c.HTML(200, "top.tmpl", nil)
}
