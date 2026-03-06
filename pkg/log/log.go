package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log 全局日志对象
// Global logger instance
var Log *zap.Logger

func L() *zap.Logger {
	if Log != nil {
		return Log
	}
	return zap.L()
}

// Init 初始化日志配置
// Init initializes the logger configuration
func Init() {
	// 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式使用 ISO8601 (例如: 2023-10-01T12:00:00.000Z)
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志级别大写显示
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 创建核心配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 使用 JSON 格式输出
		zapcore.Lock(os.Stdout),               // 输出到标准输出
		zap.InfoLevel,                         // 设置最低日志级别为 Info
	)

	// 创建 Logger 实例
	Log = zap.New(core, zap.AddCaller()) // AddCaller 会在日志中添加调用者信息 (文件名和行号)

	// 替换全局的 logger，方便直接使用 zap.L()
	zap.ReplaceGlobals(Log)
}

func Info(msg string, fields ...zap.Field)  { L().Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { L().Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field) { L().Error(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { L().Fatal(msg, fields...) }
func Debug(msg string, fields ...zap.Field) { L().Debug(msg, fields...) }
