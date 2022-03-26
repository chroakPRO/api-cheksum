package main

import (
	"flag"
	"github.com/coopersec/api-cheksum/pkg/routing"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"log"
	"time"
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
	routing.Run()
}

// getrouting will create our routing of our entire application
// this way every group of routing can be defined in their own file
// so this one won't be so messy
