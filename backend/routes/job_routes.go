package routes

import (
	"github.com/danubiobwm/recruitment-system/controllers"
	"github.com/danubiobwm/recruitment-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupJobRoutes(router *gin.Engine) {
	jobs := router.Group("/api")
	jobs.Use(middleware.AuthMiddleware())

	jobs.POST("/jobs", controllers.CreateJob)
	jobs.GET("/jobs", controllers.ListJobs)
	jobs.GET("/jobs/:id", controllers.GetJob)
	jobs.PUT("/api/jobs/:id", controllers.UpdateJob)
	jobs.DELETE("/jobs/:id", controllers.DeleteJob)
}
