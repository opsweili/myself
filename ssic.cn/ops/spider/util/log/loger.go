package log

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	outputdir = "../logdir/"
	outpath   = "spider.log"
	errpath   = "spider.err"
)

var (
	logger        *zap.Logger
	encoderConfig = zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,

		//EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		//EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},

		//EncodeCaller: zapcore.FullCallerEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,

		EncodeName: zapcore.FullNameEncoder,
	}
	encoder      = zapcore.NewJSONEncoder(encoderConfig)
	automicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	writeSyncer = zapcore.AddSync(os.Stdout)
)

func init() {

	// 初始化设置日志级别(automicLevel) debug 和输出设备(writeSyncer), stdout
	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			writeSyncer,
			automicLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1))
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		// 没有使用go风格反人类的format格式
		outputdir+filename+".%Y%m%d",
		//rotatelogs.WithLinkName(outputdir+outpath+".log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		//panic(err)
		fmt.Println("create file error!")
	}
	return hook
}

// SetLevel 设置日志级别
func SetLevel(s string) {
	if level, ok := levelMap[s]; ok {

		if s == "error" {
			automicLevel.SetLevel(level)
			writeSyncer = zapcore.AddSync(getWriter(outpath))
			core := zapcore.NewTee(
				zapcore.NewCore(
					encoder,
					writeSyncer,
					automicLevel),
			)
			logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1))
		} else {
			automicLevel.SetLevel(level)

		}
	}
}

// Field 字段
func Field(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

// Debug 级别
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info 级别
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn 级别
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error 级别
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// DPanic 级别
func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

// Panic 级别
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// Fatal 级别
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
