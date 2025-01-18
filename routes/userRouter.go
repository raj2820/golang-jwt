package routes

import (
	controller "rajauth/controllers"
	"rajauth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users",controller.GetUser())
	incommingRoutes.GET("users/:user_id",controller.GetUser())
	
}
