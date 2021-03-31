package logger

import (
	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//NewLogger 日志配置初始化
var lg *zap.Logger

//NewLogger 日志配置初始化器
func NewLogger(cfg *setting.LogConfig) error {
	writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxAge, cfg.MaxBackups, cfg.Compress)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(globals.LogSetting.Level)); err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) //替换zap包中的全局logger实例，后续其他包中只需使用zap.L()调用即可

	return nil
}

//日志格式标准
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//日志写入格式
func getLogWriter(filename string, maxSize, maxAge, maxBackups int, compress bool) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename, //日志文件名
		MaxSize:    maxSize,  //日志文件大小
		MaxAge:     maxAge,   //
		MaxBackups: maxBackups,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}
