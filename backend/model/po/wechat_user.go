package po

import "time"

type WechatUserPO struct {
	BasePO
	// openid
	Openid string `gorm:"column:openid"`
	// 昵称
	NickName string `gorm:"column:nick_name"`
	// 头像
	AvatarUrl string `gorm:"column:avatar_url"`
	// 最后登录时间
	LastLogin time.Time `gorm:"column:last_login"`
}

func (WechatUserPO) TableName() string {
	return "wechat_user"
}
