package models

type RecordSettings struct {
	StartTimecode    string `json:"start_timecode" example:""`
	EndTimecode      string `json:"end_timecode" example:""`
	AnnotationLength int    `json:"annotation_length"  example:"200"`
	ArticleLength    int    `json:"article_length" example:"1000"`
	ScreenshotTiming int    `json:"screenshot_timing" example:"3"`
}
