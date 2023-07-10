package sl

import (
	"os"
	"time"

	"github.com/devsendjin/url-shortener/internal/constants"
	"golang.org/x/exp/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func replaceLoggerAttr(groups []string, attr slog.Attr) slog.Attr {
	if attr.Key == slog.TimeKey {
		t := time.Now()
		formattedTime := t.Format(constants.FullDateTime)

		return slog.Attr{
			Key:   slog.TimeKey,
			Value: slog.StringValue(formattedTime),
		}
	}

	return attr
}

func New(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case constants.EnvLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level:       slog.LevelDebug,
				ReplaceAttr: replaceLoggerAttr,
			}),
		)
	case constants.EnvDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level:       slog.LevelDebug,
				ReplaceAttr: replaceLoggerAttr,
			}),
		)
	case constants.EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level:       slog.LevelInfo,
				ReplaceAttr: replaceLoggerAttr,
			}),
		)
	default:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level:       slog.LevelInfo,
				ReplaceAttr: replaceLoggerAttr,
			}),
		)
	}

	return logger
}
