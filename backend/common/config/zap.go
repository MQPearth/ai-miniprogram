package config

type Zap struct {
	// 级别
	Level string `mapstructure:"level" yaml:"level"`
	// 日志前缀
	Prefix string `mapstructure:"prefix" yaml:"prefix"`
	// 输出
	Format string `mapstructure:"format" yaml:"format"`
	// 日志文件夹
	Director string `mapstructure:"director" yaml:"director"`
	// 编码级
	EncodeLevel string `mapstructure:"encode-level" yaml:"encode-level"`
	// 栈名
	StacktraceKey string `mapstructure:"stacktrace-key" yaml:"stacktrace-key"`
	// 日志留存时间
	MaxAge int `mapstructure:"max-age" yaml:"max-age"`
	// 显示行
	ShowLine bool `mapstructure:"show-line" yaml:"show-line"`
	// 输出控制台
	LogInConsole bool `mapstructure:"log-in-console" yaml:"log-in-console"`
}
