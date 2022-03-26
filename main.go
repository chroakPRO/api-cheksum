package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"log"
)

var (
	router = gin.Default()
	limit  ratelimit.Limiter
	rps    = flag.Int("rps", 100, "request per second")
)

// Run will start the server
func main() {
	flag.Parse()
	ginRun(*rps)
}

func ginRun(rps int) {

	limit = ratelimit.New(rps)

	app := gin.Default()
	app.Use(leakBucket())

	app.GET("/rate", func(ctx *gin.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
	// Send request onwards
	routes.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
