package core

import (
	"encoding/json/v2"
	"errors"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/model"
)

// Initialize will set up the main system folder /**
func Initialize() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	docs := filepath.Join(home, "Documents")
	info, err := os.Stat(docs)

	if err != nil {
		log.Fatal(err)
	}

	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("Documents folder does not exist")
	}

	if !info.IsDir() {
		log.Fatal("Documents folder is not a directory")
	}

	// Create the main folder.
	files := filepath.Join(docs, "__files")
	err = os.Mkdir(files, os.ModePerm)

	if errors.Is(err, os.ErrExist) {
		slog.Debug("Directory already exists, skipping initialization.")
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	// Create the save folder.
	err = os.Mkdir(filepath.Join(files, "save"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// GetFilesDir will get the main system folder.
func getFilesDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	files := filepath.Join(home, "Documents/__files")
	info, err := os.Stat(files)

	if err != nil {
		log.Fatal(err)
	}

	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("__files folder does not exist")
	}

	if !info.IsDir() {
		log.Fatal("__files folder is not a directory")
	}

	return files
}

func WriteSave(system model.System, index int) {
	slot, err := system.GetSlot(index)
	if err != nil {
		log.Fatal(err)
	}

	saveSlot := index + 1
	saveFile := "game_" + strconv.Itoa(saveSlot) + ".json"

	data, err := json.Marshal(slot)
	if err != nil {
		log.Fatal(err)
	}

	saveFileDir := filepath.Join(getFilesDir(), "save/"+saveFile)
	err = os.WriteFile(saveFileDir, data, os.ModePerm)
}

func ReadSave(index int) *model.Slot {
	file := filepath.Join(getFilesDir(), "save/"+strconv.Itoa(index+1)+".json")
	data, err := os.ReadFile(file)

	slog.Debug("Read save file", slog.String("file", file))

	if errors.Is(err, os.ErrNotExist) {
		slog.Debug("File does not exist", slog.String("file", file))
		return nil
	}

	if err != nil {
		log.Fatal(err)
	}

	var slot model.Slot
	err = json.Unmarshal(data, &slot)
	if err != nil {
		log.Fatal(err)
	}

	return &slot
}
