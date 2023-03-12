package services

import (
	"encoding/json"

	"github.com/PongDev/PongDev_agnos_backend_internship_2023/models"
)

func LogRequestResponse(request interface{}, response interface{}) {
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return
	}
	models.DB.Create(&models.Log{
		Request:  string(requestJSON),
		Response: string(responseJSON),
	})
}
