package router

import (
	"ai-backend/api"
	"github.com/gin-gonic/gin"
)

type WechatUserRouter struct{}

func (s *WechatUserRouter) InitPublicRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("wechat").Group("user")
	wechatUserApi := api.ApiGroup.WechatUserApi
	authorityRouter.POST("testLogin", wechatUserApi.TestLogin)
	authorityRouter.POST("login", wechatUserApi.Login)
}

func (s *WechatUserRouter) InitPrivateRouter(Router *gin.RouterGroup) {

}
