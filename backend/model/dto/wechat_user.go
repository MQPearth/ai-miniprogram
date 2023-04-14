package dto

type LoginDTO struct {
	Token     string `json:"token"`
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
}

type LoginReqDTO struct {
	Openid    string
	NickName  string
	AvatarUrl string
}
