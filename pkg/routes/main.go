package routes

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"log"
	"time"
)

var (
	router = gin.Default()
)

func Run() {
	getRoutes()
	router.Run(":5000")
}

func getRoutes() {

	v1 := router.Group("/v1")
	addPersonalRoute(v1)
	addAuth(v1)
	// v2 := router.Group("/v2")

}

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

func ginRun(rps int) {

	limit = ratelimit.New(rps)

	app := gin.New()
	app.Use(leakBucket())

	app.GET("/rate", func(ctx *gin.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
	// Send request onwards
	router.Run(":5000")
}
