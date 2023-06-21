package models

// MediaCreate model info
//
//	@Description	Media create model
type MediaCreate struct {
	Link string `json:"link" binding:"required" example:"https://i.pinimg.com/736x/f4/d2/96/f4d2961b652880be432fb9580891ed62.jpg" format:"url"`
}

// MediaDb model info
//
//	@Description	Media db model
type Media struct {
	Id   int    `json:"id" binding:"required" example:"1"`
	Link string `json:"link" binding:"required" example:"https://i.pinimg.com/736x/f4/d2/96/f4d2961b652880be432fb9580891ed62.jpg" format:"url"`
}
