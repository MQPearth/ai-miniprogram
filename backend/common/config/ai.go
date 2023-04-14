package config

type Ai struct {
	// token
	Token string `mapstructure:"token" yaml:"token"`
	// 模型
	Model string `mapstructure:"model" yaml:"model"`
	// EnableProxy
	EnableProxy bool `mapstructure:"enable-proxy" yaml:"enable-proxy"`
	// proxy
	Proxy string `mapstructure:"proxy" yaml:"proxy"`
}
