package vo

type SaveFeedbackVO struct {
	// 内容
	Content string `json:"content" binding:"required,max=255"`
}
