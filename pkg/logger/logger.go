package logger

import (
	"log/slog"
	"os"

	"poc-recognition/pkg/config"
)

var Logger *slog.Logger

func Init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel(config.LOG_LEVEL),
		AddSource: true,
	}))

	slog.SetDefault(Logger)
}

func logLevel(level string) slog.Level {
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
	if l, ok := levels[level]; ok {
		return l
	}

	return slog.LevelInfo
}

func Info(msg string, args interface{}) {
	Logger.Info(msg, args)
}

func Warning(msg string, args interface{}) {
	Logger.Warn(msg, args)
}

func Error(msg string, args interface{}) {
	Logger.Error(msg, args)
}

func Debug(msg string, args interface{}) {
	Logger.Debug(msg, args)
}
