package endpoint

import (
	"github.com/coopersec/api-cheksum/app/models"
	"github.com/gin-gonic/gin"
)

func GetPersonalInfo(c *gin.Context) error {

	Info := &models.PersonalStruct{}

	// Return 200 status
	return c.JSON(200, gin.H{
		"UUID":      Info.UUID,
		"name":      Info.Age,
		"portfolio": Info.Portfolio,
		"website":   Info.Website,
		"employed":  Info.Employed,
		"email":     Info.Email,
	})
}
