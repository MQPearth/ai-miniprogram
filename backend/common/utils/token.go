package utils

import (
	"ai-backend/common/cons"
	"ai-backend/model/common/resp"
	"ai-backend/model/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
	"time"
)

type tokenUtil struct {
}

var Token = new(tokenUtil)

func (t *tokenUtil) TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(cons.RequestHeaderTokenKey)
		if token == "" {
			c.JSON(http.StatusOK, resp.Response{
				Code: resp.TokenExpired,
				Msg:  "未登录",
			})
			c.Abort()
			return
		}
		loginCache := middleware.RedisLoginCache{}
		// 取值
		redisTokenKey := t.GetRedisTokenKey(token)
		err := Redis.GetObject(redisTokenKey, &loginCache)
		if err == redis.Nil {
			c.JSON(http.StatusOK, resp.Response{
				Code: resp.TokenExpired,
				Msg:  "登录已失效",
			})
			c.Abort()
			return
		} else if err != nil {
			resp.FailWithMessage("系统错误", c)
			c.Abort()
			return
		}
		// 刷新时间
		loginCache.LastRequestTime = time.Now()
		if err := Redis.Set(redisTokenKey, loginCache, cons.TokenCacheTimeout); err != nil {
			resp.FailWithMessage("系统错误", c)
			c.Abort()
			return
		}
		c.Set(cons.RequestInfoCache, loginCache)
		c.Next()
	}
}

func (t *tokenUtil) GenerateToken() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func (t *tokenUtil) GetRedisTokenKey(token string) string {
	return cons.TokenCacheKeyPrefix + token
}

func (t *tokenUtil) GetTokenCache(c *gin.Context) middleware.RedisLoginCache {
	return c.MustGet(cons.RequestInfoCache).(middleware.RedisLoginCache)
}

func (t *tokenUtil) GetTokenCacheUserId(c *gin.Context) uint {
	cache := t.GetTokenCache(c)
	return cache.Id
}
