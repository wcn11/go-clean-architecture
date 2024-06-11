package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	// Set the output of gRPC to standard output
	Logger.SetOutput(os.Stdout)
	// Set the logger level
	Logger.SetLevel(logrus.InfoLevel)
	// Set the logger format
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
