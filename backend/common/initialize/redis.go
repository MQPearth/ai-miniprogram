package initialize

import (
	"ai-backend/common/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr: redisCfg.Url,
		// no password set
		Password: redisCfg.Password,
		// use default DB
		DB: redisCfg.Db,
	})
	pong, err := client.Ping(context.TODO()).Result()
	if err != nil {
		global.GLog.Error("Redis 连接失败, 信息:", zap.Error(err))
		panic("Redis 连接失败")
	} else {
		global.GLog.Info(fmt.Sprintf("Redis连接成功 ping: %s", pong))
		global.GRedis = client
	}
}
