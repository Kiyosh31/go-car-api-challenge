package main

import (
	"cars-go/api"
	"cars-go/config"
	"cars-go/store"
	"os"

	"log/slog"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "date"
				a.Value = slog.Int64Value(a.Value.Time().Unix())
			}
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.Int("What is ", 42),
		slog.Group("votes",
			slog.Int("pikachu", 40),
			slog.Int("mew", 22),
		),
	})

	logger := slog.New(logHandler)
	logger.Info("aqui es1")
	logger.Debug("aqui es2")
	logger.Warn("aqui es3")
	logger.Error("aqui es4")

	env, err := config.LoadEnvVars()
	if err != nil {
		// log.Er("Could not load env vars: %v", err)
	}

	store := store.NewCarStore()
	server := api.CreateNewServer(env.Port, *store)

	err = server.Start()
	if err != nil {
		// log.Fatalf("Error running server: %v", err)
	}

	// log.Printf("Server running in port: %v", env.Port)
	// slog.Info("Aqui")

}
