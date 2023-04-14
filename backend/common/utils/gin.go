package utils

import (
	"ai-backend/common/cons"
	"ai-backend/common/global"
	"ai-backend/model/common/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// RequestLimit 请求限制
// 由于要获取到当前登录的用户, 所以该处理器只能在common/utils/token#TokenAuth之后运行
func RequestLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := Token.GetTokenCacheUserId(c)

		redisKey := fmt.Sprintf(cons.RequestNamingQty+"%d", id)
		// 一天只能调用3次
		qtyStr, err := Redis.GetString(redisKey)
		if err != nil {
			resp.FailWithMessage(err.Error(), c)
			return
		}
		var qty int
		if qtyStr == "" {
			qty = 1
		} else {
			qty, _ = strconv.Atoi(qtyStr)
			if qty >= 5 {
				c.JSON(http.StatusOK, resp.Response{
					Code: resp.RequestLimit,
					Msg:  "由于服务器资源有限, 每人每天只能使用5次, 现已达到上限",
				})
				c.Abort()
				return
			}
			qty++
		}

		c.Next()

		requestOk := c.GetBool(cons.RequestProcessOk)

		if requestOk {
			// 请求成功后计算次数
			// 计算现在到明天凌晨的剩余时间
			now := time.Now()
			// 计算明天凌晨的时间
			tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			// 计算剩余时间
			remainingTime := tomorrow.Sub(now)

			err = Redis.Set(redisKey, qty, remainingTime)
			if err != nil {
				global.GLog.Error("redis set error", zap.Error(err))
				// 默认上层已设置了response body
				c.Abort()
				return
			}
		}
	}
}

func AutoBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		status := c.Writer.Status()
		banKey := cons.IpBlackList + ip
		val, err := Redis.GetString(banKey)
		if err != nil || val != "" {
			// 已被封禁
			global.GLog.Info(fmt.Sprintf("已被封禁的IP: %s", ip))
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if status == 404 {
			err := Redis.Set(banKey, time.Now(), time.Hour*24*180)
			if err != nil {
				global.GLog.Error("AutoBan", zap.Error(err))
			}
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
		return
	}
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start).Milliseconds()

		if query != "" {
			path = path + "?" + query
		}

		// 写入日志
		global.GLog.Info(fmt.Sprintf("[GIN] |%s|%9s |%17s |%s \"%s\"", colorfulStatus(c.Writer.Status()),
			fmt.Sprintf("%dms", cost),
			c.ClientIP(), colorfulMethod(c.Request.Method), path),
			zap.String("User-Agent", c.Request.UserAgent()))
	}
}

func colorfulStatus(status int) string {
	var colorCode string
	if status == 200 {
		colorCode = "\033[0;42m %d \033[0m"
	} else if status == 404 {
		colorCode = "\033[0;43;30m %d \033[0m"
	} else {
		colorCode = "\033[0;41m %d \033[0m"
	}
	return fmt.Sprintf(colorCode, status)
}

func colorfulMethod(method string) string {
	var colorCode string
	switch method {
	case "POST":

		colorCode = "\033[0;48;2;120;220;232;38;2;255;255;255m %-8s\033[0m"
	case "GET":
		colorCode = "\033[0;48;2;252;152;103;38;2;255;255;255m %-8s\033[0m"
	case "PUT":
		colorCode = "\033[0;43;30m %-8s\033[0m"
	case "DELETE":
		colorCode = "\033[0;48;2;255;97;136;97m %-8s\033[0m"
	case "OPTIONS":
		colorCode = "\033[0;48;2;147;146;147;38;2;98;88;98m %-8s\033[0m"
	default:
		colorCode = " %-8s"
	}
	return fmt.Sprintf(colorCode, method)
}

// GinRecovery recover掉项目可能出现的panic
// 此函数是捕获panic，根据gin框架内的Recovery修改的
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequestByteArr, _ := httputil.DumpRequest(c.Request, true)
				httpRequestString := strings.ReplaceAll(string(httpRequestByteArr), "\r\n", " ")

				var errInfo string
				if brokenPipe {
					errInfo = c.Request.URL.Path
				} else {
					errInfo = "[Recovery from panic]"
				}

				global.GLog.Error(errInfo, zap.Any("error", err), zap.String("request", httpRequestString))

				if brokenPipe {
					_ = c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					stack := string(debug.Stack())
					global.GLog.Error(stack)
				}

				c.JSON(http.StatusInternalServerError, resp.Response{
					Code: resp.Error,
					Msg:  "server error",
				})
			}
		}()
		c.Next()
	}
}
