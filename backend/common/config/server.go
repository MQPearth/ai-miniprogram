package config

type Server struct {
	// 服务器地址
	Port    int    `mapstructure:"port" yaml:"port"`
	GinMode string `mapstructure:"gin-mode" yaml:"gin-mode"`
}
