package router

import (
	"ai-backend/api"
	"github.com/gin-gonic/gin"
)

type TranslateRouter struct{}

func (s *TranslateRouter) InitPublicRouter(Router *gin.RouterGroup) {

}

func (s *TranslateRouter) InitLimitRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("translate")
	translateApi := api.ApiGroup.TranslateApi
	authorityRouter.POST("", translateApi.Translate)
}
