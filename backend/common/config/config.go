package config

type Config struct {
	// gorm
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
	// Zap
	Zap Zap `mapstructure:"zap" yaml:"zap"`
	// Server
	Server Server `mapstructure:"server" yaml:"server"`
	// Redis
	Redis Redis `mapstructure:"redis" yaml:"redis"`
	// ai
	Ai Ai `mapstructure:"ai" yaml:"ai"`
	// weixin
	Weixin Weixin `mapstructure:"weixin" yaml:"weixin"`
}
