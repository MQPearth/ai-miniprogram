package global

import (
	"ai-backend/common/config"
	"github.com/go-redis/redis/v8"
	ai "github.com/sashabaranov/go-openai"
	"github.com/silenceper/wechat/v2/miniprogram"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GConfig      config.Config
	GLog         *zap.Logger
	GDb          *gorm.DB
	GRedis       *redis.Client
	GAi          *ai.Client
	GMiniprogram *miniprogram.MiniProgram
)
