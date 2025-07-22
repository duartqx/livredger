package logger

import (
	"log/slog"
	"os"
)

var SLogger = slog.New(
	slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{
			Level: func() *slog.LevelVar {
				lv := slog.LevelVar{}
				lv.Set(slog.LevelDebug)
				return &lv
			}(),
		},
	),
)
