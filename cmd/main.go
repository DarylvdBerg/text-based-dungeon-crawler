package main

import (
	"log"
	"log/slog"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/core"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/ui"
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

	//system := model.System{}
	core.Initialize()
	p := tea.NewProgram(ui.New())
	if _, err := p.Run(); err != nil {
		slog.Error("tui crashed", slog.Any("err", err))
		os.Exit(1)
	}
}
