package base

import (
	"log/slog"
	"os"
)

func LogSetting() {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})
	slog.SetDefault(slog.New(handler))
}
