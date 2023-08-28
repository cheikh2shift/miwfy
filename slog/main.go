package main

import (
	"os"

	"log/slog"
)

func main() {

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// Set a custom level to show all log output. The default value is
		// LevelInfo, which would drop Debug and Trace logs.
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	})

	logger := slog.New(th)

	logger.Debug("Hello", "Who", "World")
}
