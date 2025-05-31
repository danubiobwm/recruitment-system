package routes

import (
	"github.com/danubiobwm/recruitment-system/controllers"
	"github.com/danubiobwm/recruitment-system/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/me", controllers.Me)

	return r
}
