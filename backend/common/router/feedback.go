package router

import (
	"ai-backend/api"
	"github.com/gin-gonic/gin"
)

type FeedbackRouter struct{}

func (s *FeedbackRouter) InitPublicRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("feedback")
	feedbackApi := api.ApiGroup.FeedbackApi
	authorityRouter.GET("view", feedbackApi.View)
}

func (s *FeedbackRouter) InitPrivateRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("feedback")
	feedbackApi := api.ApiGroup.FeedbackApi
	authorityRouter.POST("save", feedbackApi.Save)
}
