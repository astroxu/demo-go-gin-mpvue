package routers

import (
	"github.com/gin-gonic/gin"
	"src/controllers/api"
)

func InitRouter() {
	r := gin.Default()
	// Simple group: v1
	v1 := r.Group("api/v1")
	{
		v1.POST("/users/signin", api.SignIn)
		v1.POST("/users/register", api.Register)
		v1.GET("/users/:id", api.GetUser)
		v1.PUT("/users/:id", api.PutUser)
		v1.DELETE("/users/:id", api.DelUser)
		v1.GET("/users", api.GetUsers)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
