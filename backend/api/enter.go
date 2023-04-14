package api

import "ai-backend/service"

type ApiGroupEnter struct {
	WechatUserApi
	NamingApi
	FeedbackApi
	TranslateApi
	AdApi
}

var (
	// ApiGroup public
	ApiGroup = new(ApiGroupEnter)

	// wechatUserService private
	wechatUserService = service.ServiceGroup.WechatUserService
	// feedBackService
	feedBackService = service.ServiceGroup.FeedBackService
)
