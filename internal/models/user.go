package models

// UserBase model info
//
//	@Description	User base model
type UserBase struct {
}

// UserCreate model info
//
//	@Description	User create model
type UserCreate struct {
	Email    string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
	Password string `json:"password" binding:"required" example:"test"`
}

// UserLogin model info
//
//	@Description	User login model
type UserLogin struct {
	Email    string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
	Password string `json:"password" binding:"required" example:"test"`
}

// UserDb model info
//
//	@Description	User db model
type UserDb struct {
	Id       int
	Email    string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
	Password string
}

// UserDto model info
//
//	@Description	User dto model
type UserDto struct {
	Id    int    `json:"id" binding:"required" example:"1"`
	Email string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
}
