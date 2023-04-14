package initialize

import (
	"ai-backend/common/global"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Viper() *viper.Viper {
	v := viper.New()
	var env string
	args := os.Args
	if len(args) == 1 {
		env = "dev"
	} else {
		env = os.Args[1]
	}
	fmt.Printf("当前环境为: %s\n", env)
	v.SetConfigFile("config-" + env + ".yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}
	// 解析 配置文件的值注入结构体
	if err = v.Unmarshal(&global.GConfig); err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}

	return v
}
