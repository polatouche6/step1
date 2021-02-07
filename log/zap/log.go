package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)
var (
	sugarLogger *zap.SugaredLogger
	logger *zap.Logger
	err error
	core zapcore.Core
)

func getAbsolutePath(fpath string) (string, error) {
	abs := filepath.IsAbs(fpath)
	if abs {
		return fpath, nil
	}
	_, fpath, _, _ = runtime.Caller(1)
	return fpath, nil
}


func InitLogger(filePath string) error {
	filePath, err := getAbsolutePath(filePath)
	if err != nil {
		return err
	}
	writeSyncer := getLogWriter(filePath)
	encoder := getEncoder()
	core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// First, define our level-handling logic.
	// 仅打印Error级别以上的日志
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// 打印所有级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	hook := lumberjack.Logger{
		Filename:   "/tmp/abc.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 3,
		MaxAge:     7,    // days
		Compress:   true, // disabled by default
	}


	topicErrors := zapcore.AddSync(ioutil.Discard)
	fileWriter := zapcore.AddSync(&hook)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core = zapcore.NewTee(
		// 打印在kafka topic中（伪造的case）
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		// 打印在控制台
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		// 打印在文件中
		zapcore.NewCore(consoleEncoder, fileWriter, highPriority),
	)

	return nil
}
// https://www.cnblogs.com/feixiangmanon/p/11109174.html
//func initCfg() error {
//	cfg := zap.Config{
//		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
//		Development: true,
//		Encoding:    "json",
//		Sampling:    zap.SamplingConfig{
//			Initial: 1,
//			Thereafter: 1,
//			Hook:
//		},
//		EncoderConfig: zapcore.EncoderConfig{
//			TimeKey:        "t",
//			LevelKey:       "level",
//			NameKey:        "logger",
//			CallerKey:      "caller",
//			MessageKey:     "msg",
//			StacktraceKey:  "trace",
//			LineEnding:     zapcore.DefaultLineEnding,
//			EncodeLevel:    zapcore.LowercaseLevelEncoder,
//			EncodeTime:     formatEncodeTime,
//			EncodeDuration: zapcore.SecondsDurationEncoder,
//			EncodeCaller:   zapcore.ShortCallerEncoder,
//		},
//		OutputPaths:      []string{"/tmp/zap.log"},
//		ErrorOutputPaths: []string{"/tmp/zap.log"},
//		InitialFields: map[string]interface{}{
//			"app": "test",
//		},
//	}
//	logger, err = cfg.Build()
//	if err != nil {
//		return err
//	}
//	logger.Error()
//	return nil
//
//}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(filePath string) zapcore.WriteSyncer {
	file, _ := os.Create(filePath)
	return zapcore.AddSync(file)
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func Hook(zapcore.Entry, zapcore.SamplingDecision) {

}

func WithConsoleLog()  {

}

func WithFileLog(){

}

func WithHook() {

}