package logger

import (
	"despair/app"
	"github.com/wobusbz/logger"
)

func InitLogger() {
	l := logger.NewLogger()
	l.FileName = app.App.Name
	l.FilePath = app.App.RootDir
	l.SetLoggerMax(300)
	l.SetConsole(false)
	logger.CustomLogger(l)
}

func Debug(format string, args ...interface{}) {
	logger.Debug(format, args)
}

func Info(format string, args ...interface{}) {
	logger.Info(format, args)
}

func Warn(format string, args ...interface{}) {
	logger.Warn(format, args)
}

func Error(format string, args ...interface{}) {
	logger.Error(format, args)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatal(format, args)
}
