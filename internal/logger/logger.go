package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// AppLogger is the global logger
var AppLogger = logrus.New()

// InitLogger configures logging to both console and a rotating file
func InitLogger(logFile string, level logrus.Level) {
	// 1. Configure log rotation
	logFileWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // MB per log file
		MaxBackups: 5,  // keep last 5 old files
		MaxAge:     30, // days
		Compress:   true,
	}

	// 2. Combine stdout and log file writers
	multiWriter := io.MultiWriter(os.Stdout, logFileWriter)

	// 3. Configure the logger
	AppLogger.SetOutput(multiWriter)
	AppLogger.SetLevel(level)                       // DebugLevel, InfoLevel, WarnLevel, ErrorLevel
	AppLogger.SetFormatter(&logrus.JSONFormatter{}) // JSON is better for production
}
