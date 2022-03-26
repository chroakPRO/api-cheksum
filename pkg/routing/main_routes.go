package routing

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
	addPersonalrouting(v1)
	addAuthrouting(v1)

	// v2 := router.Group("/v2")

}
