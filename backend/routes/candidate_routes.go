package routes

import (
	"github.com/danubiobwm/recruitment-system/controllers"
	"github.com/danubiobwm/recruitment-system/middleware"
	"github.com/gin-gonic/gin"
)

func SetupCandidateRoutes(r *gin.Engine) {
	candidates := r.Group("/api/candidates")
	candidates.Use(middleware.AuthMiddleware())
	{
		candidates.POST("", controllers.CreateCandidate)
		candidates.GET("", controllers.GetCandidates)
		candidates.GET(":id", controllers.GetCandidate)
		candidates.PUT(":id", controllers.UpdateCandidate)
		candidates.DELETE(":id", controllers.DeleteCandidate)
	}
}
