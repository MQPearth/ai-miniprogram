package vo

type TranslateVO struct {
	// 目标语言
	Target string `json:"target" binding:"required,max=20"`
	// 内容
	Content string `json:"content" binding:"required,max=1000"`
}
