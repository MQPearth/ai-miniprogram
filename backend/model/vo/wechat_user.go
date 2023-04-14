package vo

type TestLoginVO struct {
	// openid
	Openid    string `json:"openid" binding:"required"`
	NickName  string `json:"nickName" binding:"required"`
	AvatarUrl string `json:"avatarUrl" binding:"required"`
}

type LoginVO struct {
	// code
	Code string `json:"code" binding:"required"`
	// appid
	Appid string `json:"appid" binding:"required"`

	NickName  string `json:"nickName" binding:"required"`
	AvatarUrl string `json:"avatarUrl" binding:"required"`
}
