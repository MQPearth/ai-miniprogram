package api

import (
	"ai-backend/common/global"
	"ai-backend/common/utils"
	"ai-backend/model/common/resp"
	"ai-backend/model/dto"
	"ai-backend/model/vo"
	"github.com/gin-gonic/gin"
)

type FeedbackApi struct{}

// Save 保持反馈
func (a *FeedbackApi) Save(c *gin.Context) {
	feedbackVo := vo.SaveFeedbackVO{}
	err := c.ShouldBindJSON(&feedbackVo)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	feedbackDto := dto.SaveFeedbackDTO{Content: feedbackVo.Content, UserId: utils.Token.GetTokenCacheUserId(c)}
	err = feedBackService.SaveFeedback(&feedbackDto)
	if err != nil {
		resp.Fail(c)
		return
	}
	resp.Ok(c)
}

// View 查看首页
func (a *FeedbackApi) View(c *gin.Context) {
	resp.OkWithData(global.GConfig.Weixin.Version, c)
}
