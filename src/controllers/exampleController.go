package controllers

import (
	"go-template-rest-api/src/responses/exampleresponse"
	"go-template-rest-api/src/structs/examplestruct"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// ExampleController struct
type ExampleController struct{}

// Index function
func (e *ExampleController) Index(c *gin.Context) {
	exampleresponse.ExampleIndexResponse(c)
}

// Create function
func (e *ExampleController) Create(c *gin.Context) {

	var newMyGuests examplestruct.MyGuests
	c.Bind(&newMyGuests)

	_, err := govalidator.ValidateStruct(newMyGuests)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err.Error(),
		})
		c.Abort()
		return
	}

	exampleresponse.ExampleCreateResponse(c, newMyGuests)
}
