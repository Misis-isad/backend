package models

// UserCreate model info
//
//	@Description	User create model
type UserCreate struct {
	Email    string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
	Password string `json:"password" binding:"required" example:"test"`
	Fio      string `json:"fio" binding:"required" example:"Мисосов Мисос Мисосович"`
}

// UserLogin model info
//
//	@Description	User login model
type UserLogin struct {
	Email    string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
	Password string `json:"password" binding:"required" example:"test"`
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique; index"`
	Password string
	Fio      string
	Records  []Record  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// UserDto model info
//
//	@Description	User dto model
type UserDto struct {
	ID    uint   `json:"id" binding:"required" example:"1"`
	Email string `json:"email" binding:"required" format:"email" example:"test@test.ru"`
	Fio   string `json:"fio" binding:"required" example:"Мисосов Мисос Мисосович"`
}
