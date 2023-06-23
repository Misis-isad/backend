package models

// TokenResponse model info
//
//	@Description	Token response model
type TokenResponse struct {
	Token     string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	TokenType string `json:"token_type" binding:"required" example:"Bearer"`
}

type MlResponse struct {
	Body           string `json:"body" binding:"required" example:"{html page}" format:"html"`
	Title          string `json:"title" binding:"required" example:"{title}" format:"string"`
	PreviewPicture string `json:"preview_picture" binding:"required" example:"{url}" format:"url"`
	// MediaUrls	  []string `json:"media_urls" binding:"required" example:"[{url}]" format:"url"`
}
