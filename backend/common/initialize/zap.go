package initialize

import (
	"ai-backend/common/global"
	"ai-backend/common/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"time"
)

type Writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *Writer {
	return &Writer{Writer: w}
}

// Printf 格式化打印日志
func (w *Writer) Printf(message string, data ...interface{}) {
	if global.GConfig.Mysql.LogZap {
		global.GLog.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func ZapEncodeLevel() zapcore.LevelEncoder {
	encodeLevel := global.GConfig.Zap.EncodeLevel
	switch encodeLevel {
	case "LowercaseLevelEncoder":
		// 小写编码器
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		// 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		// 大写编码器
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		// 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
func TransportLevel() zapcore.Level {
	level := global.GConfig.Zap.Level
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GConfig.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

func GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GConfig.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    ZapEncodeLevel(),
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func GetEncoder() zapcore.Encoder {
	if global.GConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(GetEncoderConfig())
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := utils.FileRotatelogs.GetWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}

	return zapcore.NewCore(GetEncoder(), writer, level)
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

func GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, GetEncoderCore(level, GetLevelPriority(level)))
	}
	return cores
}

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.GConfig.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GConfig.Zap.Director)
		_ = os.Mkdir(global.GConfig.Zap.Director, os.ModePerm)
	}

	cores := GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
