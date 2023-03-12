package types

type PasswordRecommendationStepRequestDTO struct {
	InitPassword string `json:"init_password"`
}

type PasswordRecommendationStepResponseDTO struct {
	NumOfSteps int `json:"num_of_steps"`
}
