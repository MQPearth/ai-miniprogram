package config

type Redis struct {
	// redis的哪个数据库
	Db int `mapstructure:"db"  yaml:"db"`
	// 服务器地址:端口
	Url string `mapstructure:"url"  yaml:"url"`
	// 密码
	Password string `mapstructure:"password" yaml:"password"`
}
