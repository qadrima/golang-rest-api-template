package exampleresponse

import (
	"go-template-rest-api/src/models"
	"go-template-rest-api/src/structs/examplestruct"

	"github.com/gin-gonic/gin"
)

var exampleModel = new(models.ExampleModel)

// ExampleIndexResponse :
func ExampleIndexResponse(c *gin.Context) {

	data, err := exampleModel.GetAll()
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	var response examplestruct.MyGuestsIndexResponse
	response.Status = 200
	response.Data = data

	c.JSON(200, response)
}
