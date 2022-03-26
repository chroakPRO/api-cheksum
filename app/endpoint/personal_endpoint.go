package endpoint

import (
	"github.com/coopersec/api-cheksum/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPersonalInfo(c *gin.Context) {

	Info := &models.PersonalStruct{}

	// Return 200 status
	c.JSON(http.StatusOK, gin.H{
		"UUID":      Info.UUID,
		"name":      Info.Age,
		"portfolio": Info.Portfolio,
		"website":   Info.Website,
		"employed":  Info.Employed,
		"email":     Info.Email,
	})
}
