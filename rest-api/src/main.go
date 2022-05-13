package main

import (
	"github.com/coopersec/api-cheksum/rest-api/src/pkg/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/api"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// gin-swagger middleware
// swagger embed files
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
	r := gin.New()

	r.GET("/v2/testapi/get-string-by-int/:some_id", api.GetStringByInt)
	r.GET("/v2/testapi/get-struct-array-by-string/:some_id", api.GetStructArrayByString)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	routes.Run()
	// Comment out the following line to disable swagger UI
}
