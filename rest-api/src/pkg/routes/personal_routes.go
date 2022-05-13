package routes

import (
	"github.com/coopersec/api-cheksum/rest-api/src/app/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func addPersonalRoute(rg *gin.RouterGroup) {
	users := rg.Group("/info")
	users.GET("/", func(c *gin.Context) {
		id := uuid.New()
		info := &models.PersonalStruct{id, "Christopher Ek", 24,"https://cheksum.dev", "https://cheksum.dev", true, "christophereek97@gmail.com"}

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
	})

}
