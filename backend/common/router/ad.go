package router

import (
	"ai-backend/api"
	"github.com/gin-gonic/gin"
)

type AdRouter struct{}

func (s *AdRouter) InitPublicRouter(Router *gin.RouterGroup) {

}

func (s *AdRouter) InitLimitRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("ad")
	adApi := api.ApiGroup.AdApi
	authorityRouter.POST("", adApi.Ad)
}
