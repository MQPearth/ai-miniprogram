package po

type FeedbackPO struct {
	BasePO
	// 用户id
	UserId uint `gorm:"column:user_id"`
	// 反馈内容
	Content string `gorm:"column:content"`
}

func (FeedbackPO) TableName() string {
	return "feedback"
}
