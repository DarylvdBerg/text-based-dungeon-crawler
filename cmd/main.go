package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/DarylvdBerg/text-based-dungeon-crawler/system"
)

func setupLogging() (*os.File, error) {
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}

	level := slog.LevelInfo
	if os.Getenv("DEBUG") != "" {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
	return f, nil
}

func main() {
	f, err := setupLogging()
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Initialize the main folders we'll need to store game information
	system.Initialize()
}
