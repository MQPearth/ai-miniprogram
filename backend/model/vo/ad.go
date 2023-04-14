package vo

type AdVO struct {
	// 内容
	Content string `json:"content" binding:"required,max=50"`
}
