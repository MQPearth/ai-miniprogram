package initialize

import (
	"ai-backend/common/global"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

func Ai() {
	config := openai.DefaultConfig(global.GConfig.Ai.Token)
	if global.GConfig.Ai.EnableProxy {
		proxyUrl, err := url.Parse(global.GConfig.Ai.Proxy)
		if err != nil {
			panic(err)
		}
		config.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}
	global.GAi = openai.NewClientWithConfig(config)
	global.GLog.Info("AI初始化完成")
}
