package log

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

// 初始化配置
func InitZapLog() error {
	logBasePath := viper.GetString("general.log_path")
	mkdirErr := os.MkdirAll(logBasePath, 0766)
	if mkdirErr != nil {
		Error("日志目录创建失败", mkdirErr)
	}

	infoPath := logBasePath + "info.log"
	errPath := logBasePath + "error.log"
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{infoPath, "stdout"},
		ErrorOutputPaths: []string{errPath, "stderr"},
		InitialFields: map[string]interface{}{
			"app": viper.GetString("general.app_name"),
		},
	}
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		return err
	}
	return nil
}

func Info(msg string, args ...interface{}) {
	FormatLog(args).Sugar().Info(msg)
}

func Debug(msg string, args ...interface{}) {
	FormatLog(args).Sugar().Debugf(msg)
}

func Error(msg string, args ...interface{}) {
	FormatLog(args).Sugar().Error(msg)
}

func FormatLog(args []interface{}) *zap.Logger {
	log := Logger.With(ToJsonData(args))
	return log
}

func ToJsonData(args []interface{}) zap.Field {
	det := make([]string, 0)
	if len(args) > 0 {
		for _, v := range args {
			det = append(det, fmt.Sprintf("%+v", v))
		}
	}
	return zap.Any("data", det)
}
