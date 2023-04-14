package service

type ServiceGroupEnter struct {
	WechatUserService
	FeedBackService
}

var (
	ServiceGroup = new(ServiceGroupEnter)
)
