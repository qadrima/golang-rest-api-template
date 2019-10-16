package exampleresponse

import (
	"go-template-rest-api/src/structs/examplestruct"

	"github.com/gin-gonic/gin"
)

// ExampleCreateResponse :
func ExampleCreateResponse(c *gin.Context, newMyGuests examplestruct.MyGuests) {

	data, err := exampleModel.Create(newMyGuests)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	var response examplestruct.MyGuestsCreateResponse
	response.Status = 201
	response.Data = data

	c.JSON(201, response)
}

// ErrorResponse :
func ErrorResponse(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"status": 500,
		"error":  err,
	})
	c.Abort()
}
