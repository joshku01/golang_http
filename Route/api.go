package Route

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"golang_http/controller"
	_ "golang_http/docs"
)

func Router(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.GET("/ping", controller.ApiRoute)
	}

}
