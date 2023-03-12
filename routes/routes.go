package routes

import (
	"github.com/PongDev/PongDev_agnos_backend_internship_2023/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	apiRouter := APIRouter{
		service: &services.APIService{},
	}

	apiRouterGroup := r.Group("/api")
	{
		apiRouterGroup.GET("/strong_password_steps", apiRouter.PostPasswordRecommendationStep())
	}
}
