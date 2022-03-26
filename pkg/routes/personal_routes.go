package routing

import (
	"github.com/coopersec/api-cheksum/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addPersonalrouting(rg *gin.RouterGroup) {
	users := rg.Group("/info")
	users.GET("/ping", func(c *gin.Context) {
		GetPersonalInfo(c)
	})

}

func GetPersonalInfo(c *gin.Context) {

	info := &models.PersonalStruct{}

	if info != nil {
		c.JSON(http.StatusOK, gin.H{
			"UUID":      info.UUID,
			"name":      info.Age,
			"portfolio": info.Portfolio,
			"website":   info.Website,
			"employed":  info.Employed,
			"email":     info.Email,
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": "bad"})
	}
}
