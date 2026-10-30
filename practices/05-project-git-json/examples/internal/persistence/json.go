package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/go-practice-01/internal/tasks"
)

// Load treats an absent file as a new, empty store.
// Existing but malformed files are errors: they must not be overwritten.
func Load(path string) (tasks.State, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return tasks.State{NextID: 1}, nil
	}
	if err != nil {
		return tasks.State{}, fmt.Errorf("чтение %s: %w", path, err)
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	var state tasks.State
	if err := dec.Decode(&state); err != nil {
		return tasks.State{}, fmt.Errorf("разбор JSON %s: %w", path, err)
	}
	// A JSON file must contain exactly one value.
	var extra any
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		return tasks.State{}, fmt.Errorf("лишние данные после JSON в %s", path)
	}

	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return tasks.State{}, fmt.Errorf("проверка данных %s: %w", path, err)
	}
	return store.Snapshot(), nil
}

// Save writes to a temporary file and then replaces the destination.
// This is a small teaching example, not a transactional database.
func Save(path string, state tasks.State) error {
	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return fmt.Errorf("отказ сохранить некорректное состояние: %w", err)
	}
	data, err := json.MarshalIndent(store.Snapshot(), "", "  ")
	if err != nil {
		return fmt.Errorf("кодирование JSON: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".tasks-*.tmp")
	if err != nil {
		return fmt.Errorf("временный файл: %w", err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("запись JSON: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("синхронизация файла: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("закрытие файла: %w", err)
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return fmt.Errorf("замена %s: %w", path, err)
	}
	return nil
}
