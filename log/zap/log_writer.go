package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
)

type LogWriter  zapcore.Core

func SetLevel() {

}

func NewConsoleWrite(logLevel string, enableChangeLevel bool) (LogWriter, *zap.AtomicLevel) {
	consoleWrite := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	alevel := zap.NewAtomicLevel()
	http.HandleFunc("/log/console/level", alevel.ServeHTTP)
	alevel.SetLevel(zap.DebugLevel)
	console := zapcore.NewCore(consoleEncoder, consoleWrite, alevel)
	if enableChangeLevel {
		return console, &alevel
	}
	return console, nil
}

func NewFileWriter(filePath string, enableChangeLevel, splitLog bool) (LogWriter, *zap.AtomicLevel, error){
	fileWriter := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    1024, // M
		MaxBackups: 3,
		MaxAge:     7,    // days
		Compress:   true, // disabled by default
	}
	if !splitLog {

		mode := os.FileMode(0644)
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			panic()
		}
	}
	logFileWriter := zapcore.AddSync(&fileWriter)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	alevel := zap.NewAtomicLevel()
	fileCore := zapcore.NewCore(consoleEncoder, logFileWriter, alevel)
	if enableChangeLevel {
		return fileCore, &alevel
	}
	return fileCore, nil
}




