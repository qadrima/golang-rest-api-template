package routes

import (
	"go-template-rest-api/src/controllers"

	"github.com/gin-gonic/gin"
)

// DefaultRoute struct
type DefaultRoute struct{}

// Controllers initialize
var exampleController = new(controllers.ExampleController)

// InitRoute :
func (d *DefaultRoute) InitRoute(router *gin.Engine) {

	// Routes list
	router.GET("/", exampleController.Index)
	router.POST("/", exampleController.Create)

	// Wrong routes handle
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "url not found",
		})
	})
}
