package main

import (
	"flag"
	"github.com/coopersec/api-cheksum/pkg/routes"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var ()

// Run will start the server
func main() {
	routes.Run()
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

// getrouting will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
