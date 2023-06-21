package models

// TokenResponse model info
//
//	@Description	Token response model
type TokenResponse struct {
	Token    string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	TokeType string `json:"token_type" binding:"required" example:"Bearer"`
}
