package main

import (
	"github.com/mntky/foruka-go/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return 
		}

		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})
	router,Run(":8080")
}
