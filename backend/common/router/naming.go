package router

import (
	"ai-backend/api"
	"github.com/gin-gonic/gin"
)

type NamingRouter struct{}

func (s *NamingRouter) InitPublicRouter(Router *gin.RouterGroup) {

}

func (s *NamingRouter) InitLimitRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("naming")
	namingApi := api.ApiGroup.NamingApi
	authorityRouter.GET("getNameList", namingApi.GetNameList)
}
