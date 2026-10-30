package persistence

import "example.com/go-practice-01/internal/tasks"

// TODO: if path is absent, return an empty state with NextID=1.
// If the file exists, decode exactly one JSON document, validate it via Restore,
// and return a meaningful error for corrupted data. Do NOT silently reset it.
func Load(path string) (tasks.State, error) {
	return tasks.State{NextID: 1}, nil
}

// TODO: validate the state, marshal it using encoding/json, save to a
// temporary file in the same directory and replace the target using os.Rename.
// Do not remove the old file before a successful temporary write.
func Save(path string, state tasks.State) error {
	return nil
}
