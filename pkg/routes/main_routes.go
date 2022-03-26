package routes

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {
	setupRouter()
	router.Run(":5000")
}

func setupRouter() {

	v1 := router.Group("/v1")
	addPersonalRoutes(v1)
	addAuthRoutes(v1)
	addAuthRoutes(v1)

	// v2 := router.Group("/v2")

}
