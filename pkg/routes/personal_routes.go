package routes

import (
	"github.com/coopersec/api-cheksum/app/endpoint"
	"github.com/gin-gonic/gin"
)

func addPersonalRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/info")

	users.GET("/", func(c *gin.Context) {
		endpoint.GetPersonalInfo(c)
	})
}
