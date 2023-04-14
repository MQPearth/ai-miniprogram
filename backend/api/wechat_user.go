package api

import (
	"ai-backend/common/global"
	"ai-backend/model/common/resp"
	d "ai-backend/model/dto"
	v "ai-backend/model/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WechatUserApi struct{}

// Login 登录
func (a *WechatUserApi) Login(c *gin.Context) {
	var vo v.LoginVO
	err := c.ShouldBindJSON(&vo)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	if vo.Appid != global.GConfig.Weixin.Appid {
		resp.FailWithMessage("appid错误", c)
		return
	}

	session, err := global.GMiniprogram.GetAuth().Code2Session(vo.Code)
	if err != nil {
		global.GLog.Error("wechat login error", zap.Error(err))
		resp.Fail(c)
		return
	}

	dto := d.LoginReqDTO{Openid: session.OpenID, AvatarUrl: vo.AvatarUrl, NickName: vo.NickName}
	login, err := wechatUserService.Login(dto)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	resp.OkWithData(login, c)
}

// TestLogin 测试登录
func (a *WechatUserApi) TestLogin(c *gin.Context) {
	var vo v.TestLoginVO
	err := c.ShouldBindJSON(&vo)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	dto := d.LoginReqDTO{Openid: vo.Openid, AvatarUrl: vo.AvatarUrl, NickName: vo.NickName}
	login, err := wechatUserService.Login(dto)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	resp.OkWithData(login, c)
}
