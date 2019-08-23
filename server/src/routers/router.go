package routers

import (
	"github.com/gin-gonic/gin"
	"src/controllers/api"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/api/v1/Users/:id", api.GetUser)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run(":8080")
}
