package initialize

import (
	"ai-backend/common/global"
	"ai-backend/common/router"
	"ai-backend/common/utils"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(global.GConfig.Server.GinMode)
	ginDefault := gin.New()
	ginDefault.Use(utils.GinLogger(), utils.GinRecovery(true), utils.AutoBan())

	prefix := ginDefault.Group("api")

	// 不登录就能调用的接口
	publicGroup := prefix.Group("")
	{
		router.RouterGroupApp.WechatUserRouter.InitPublicRouter(publicGroup)
		router.RouterGroupApp.FeedbackRouter.InitPublicRouter(publicGroup)
	}
	// 登录后才能调用的接口
	privateGroup := prefix.Group("")
	privateGroup.Use(utils.Token.TokenAuth())
	{
		router.RouterGroupApp.WechatUserRouter.InitPrivateRouter(privateGroup)
		router.RouterGroupApp.FeedbackRouter.InitPrivateRouter(privateGroup)
	}
	// 登录且有请求次数限制的接口
	requestLimitGroup := prefix.Group("")
	requestLimitGroup.Use(utils.Token.TokenAuth()).Use(utils.RequestLimit())
	{
		router.RouterGroupApp.NamingRouter.InitLimitRouter(requestLimitGroup)
		router.RouterGroupApp.TranslateRouter.InitLimitRouter(requestLimitGroup)
		router.RouterGroupApp.AdRouter.InitLimitRouter(requestLimitGroup)
	}

	global.GLog.Info("路由注册完成")
	return ginDefault
}
