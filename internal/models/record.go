package models

// RecordCreate model info
//
//	@Description	Record create model
type RecordCreate struct {
	VideoLink string         `json:"video_link" binding:"required" example:"https://www.youtube.com/watch?v=4O3UGW-Bbbw" format:"url"`
	Settings  RecordSettings `json:"settings" binding:"required"`
}

type Record struct {
	ID             uint `gorm:"primaryKey"`
	Title          string
	VideoLink      string
	Published      bool `gorm:"default:false"`
	PreviewPicture string
	*RecordSettings
	UserID   uint
	Article  Article   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (r *Record) ToDto() RecordDto {
	return RecordDto{
		ID:             r.ID,
		Title:          r.Title,
		VideoLink:      r.VideoLink,
		Published:      r.Published,
		PreviewPicture: r.PreviewPicture,
	}
}

// RecordDto model info
//
//	@Description	Record dto model
type RecordDto struct {
	ID             uint   `json:"id" binding:"required" example:"1"`
	Title          string `json:"title" binding:"required" example:"title"`
	VideoLink      string `json:"video_link" binding:"required" example:"https://www.youtube.com/watch?v=4O3UGW-Bbbw" format:"url"`
	Published      bool   `json:"published" binding:"required" example:"false"`
	PreviewPicture string `json:"preview_picture" binding:"required" example:"picture url" format:"url"`
}
