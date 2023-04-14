package api

import (
	"ai-backend/common/cons"
	"ai-backend/common/global"
	"ai-backend/common/utils"
	"ai-backend/model/common/resp"
	v "ai-backend/model/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdApi struct{}

// Ad Ad
func (a *AdApi) Ad(c *gin.Context) {
	var vo v.AdVO
	err := c.ShouldBindJSON(&vo)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	question := fmt.Sprintf("帮我把这段以$开头, 以^结束的产品介绍想一段文案, 只需要回答广告文案内容, 不需要回答其他内容, 用中文回答我, 并必须将$和^去除\n $%s^",
		vo.Content)
	global.GLog.Info("question: ", zap.String("question", question))
	answer, err := utils.GetAnswer(question)
	if err != nil {
		resp.FailWithMessage("请求失败", c)
		return
	}
	c.Set(cons.RequestProcessOk, true)
	resp.OkWithData(answer, c)
}
