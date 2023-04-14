package router

type RouterGroup struct {
	WechatUserRouter WechatUserRouter
	NamingRouter     NamingRouter
	FeedbackRouter   FeedbackRouter
	TranslateRouter  TranslateRouter
	AdRouter         AdRouter
}

var RouterGroupApp = new(RouterGroup)
