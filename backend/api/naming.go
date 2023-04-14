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

type NamingApi struct{}

// GetNameList 取名
func (a *NamingApi) GetNameList(c *gin.Context) {

	var vo v.NamingVO
	err := c.ShouldBindQuery(&vo)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	question := fmt.Sprintf("帮我生成%d个中国%s生名字, 姓名必须是%s风格的, 姓氏为%s, 而且是%d个字的, 并解释意义,"+
		" 用中文回答我, 你不需要提问",
		vo.Qty, vo.Gender, vo.Style, vo.Surname, vo.Length)
	global.GLog.Info("question: ", zap.String("question", question))
	answer, err := utils.GetAnswer(question)
	if err != nil {
		resp.FailWithMessage("请求失败", c)
		return
	}
	c.Set(cons.RequestProcessOk, true)
	resp.OkWithData(answer, c)
}
