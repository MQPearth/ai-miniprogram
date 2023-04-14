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

type TranslateApi struct{}

// Translate 翻译
func (a *TranslateApi) Translate(c *gin.Context) {

	var vo v.TranslateVO
	err := c.ShouldBindJSON(&vo)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	question := fmt.Sprintf("帮我把这段以$开头, 以^结束的这段话翻译成%s, 只需要回答翻译内容, 不需要回答其他内容, 并必须将$和^去除\n $%s^",
		vo.Target, vo.Content)
	global.GLog.Info("question: ", zap.String("question", question))
	answer, err := utils.GetAnswer(question)
	if err != nil {
		resp.FailWithMessage("请求失败", c)
		return
	}
	c.Set(cons.RequestProcessOk, true)
	resp.OkWithData(answer, c)
}
