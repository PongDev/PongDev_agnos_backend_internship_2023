package routes

import (
	"net/http"

	"github.com/PongDev/PongDev_agnos_backend_internship_2023/services"
	"github.com/PongDev/PongDev_agnos_backend_internship_2023/types"
	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	service *services.APIService
}

func (a *APIRouter) PostPasswordRecommendationStep() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &types.PasswordRecommendationStepRequestDTO{}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := &types.PasswordRecommendationStepResponseDTO{
			NumOfSteps: a.service.FindMinimuRecommendationmSteps(request.InitPassword),
		}

		c.JSON(http.StatusOK, response)
	}
}
