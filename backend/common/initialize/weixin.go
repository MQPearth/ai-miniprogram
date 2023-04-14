package initialize

import (
	"ai-backend/common/global"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

func Weixin() {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     global.GConfig.Weixin.Appid,
		AppSecret: global.GConfig.Weixin.Secret,
		Cache:     memory,
	}
	global.GMiniprogram = wc.GetMiniProgram(cfg)
}
