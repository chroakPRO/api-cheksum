package main

import (
	"github.com/coopersec/api-cheksum/pkg/routes"
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files" // swagger embed files
// Run will start the server
// @title           [cheksum256] Ice Trails
// @version         0.5v
// @description     many mistakes but much laughter
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1/

// @securityDefinitions.basic  BasicAuth
func main() {
	routes.Run()
	// Comment out the following line to disable swagger UI
}
