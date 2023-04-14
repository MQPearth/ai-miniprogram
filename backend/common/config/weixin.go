package config

type Weixin struct {
	Appid string `mapstructure:"appid"  yaml:"appid"`

	Secret string `mapstructure:"secret" yaml:"secret"`

	Version string `mapstructure:"version" yaml:"version"`
}
