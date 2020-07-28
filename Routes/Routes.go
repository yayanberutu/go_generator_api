package routes

import (
	"crud-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/generate")
	{
		grpl.GET("newservice/{dbname}", controllers.CreateService)
		grpl.GET("addservice/{dbname}", controllers.AddService)
	}
	return r
}
