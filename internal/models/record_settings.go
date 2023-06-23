package models

type RecordSettings struct {
	StartTimecode    string `json:"start_timecode" binding:"required" example:"00:00:00"`
	EndTimecode      string `json:"end_timecode" binding:"required" example:"00:10:00"`
	AnnotationLength int    `json:"annotation_length" binding:"required" example:"200"`
	ArticleLength    int    `json:"article_length" binding:"required" example:"1000"`
	ScreenshotTiming int    `json:"screenshot_timing" binding:"required" example:"3"`
}
