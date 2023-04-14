package main

import (
	"ai-backend/common/global"
	"ai-backend/common/initialize"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 读取配置
	initialize.Viper()
	// 初始化日志
	global.GLog = initialize.Zap()
	zap.ReplaceGlobals(global.GLog)
	// 初始化数据库链接以及DAO
	global.GDb = initialize.GormMysql()

	// 初始化redis
	initialize.Redis()

	// 初始化chatGpt
	initialize.Ai()

	initialize.Weixin()

	if global.GDb != nil {
		db, _ := global.GDb.DB()
		// 程序结束前关闭数据库链接
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.GLog.Error("数据库连接关闭失败")
			}
		}(db)
	}

	// 初始化路由
	routers := initialize.Routers()

	port := fmt.Sprintf(":%d", global.GConfig.Server.Port)

	apiServer := initialize.InitServer(port, routers)

	global.GLog.Info(fmt.Sprintf("服务成功启动在 Port: %d", global.GConfig.Server.Port))

	global.GLog.Error(apiServer.ListenAndServe().Error())
}
