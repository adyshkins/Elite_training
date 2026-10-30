package persistence

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"example.com/go-practice-01/internal/tasks"
)

func TestRoundTripAndNextID(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	store := tasks.NewMemoryStore()
	first, err := store.Add("Изучить Go", 2)
	if err != nil {
		t.Fatal(err)
	}
	_, _ = store.Add("Написать API", 3)
	if err := store.Complete(first.ID); err != nil {
		t.Fatal(err)
	}
	if err := store.Delete(first.ID); err != nil {
		t.Fatal(err)
	}
	if err := Save(path, store.Snapshot()); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	restored := tasks.NewMemoryStore()
	if err := restored.Restore(got); err != nil {
		t.Fatal(err)
	}
	if len(restored.List()) != 1 || restored.List()[0].Title != "Написать API" {
		t.Fatalf("unexpected tasks: %+v", restored.List())
	}
	newTask, err := restored.Add("Третья", 1)
	if err != nil || newTask.ID != 3 {
		t.Fatalf("ID after restart: %+v, %v", newTask, err)
	}
}

func TestLoadScenarios(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{"empty file", "", true},
		{"invalid json", "{oops", true},
		{"trailing value", "{} {}", true},
		{"unknown field", "{\"next_id\":1,\"tasks\":[],\"secret\":1}", true},
		{"invalid id counter", "{\"next_id\":0,\"tasks\":[]}", true},
		{"duplicate id", "{\"next_id\":3,\"tasks\":[{\"id\":1,\"title\":\"A\",\"status\":\"new\",\"priority\":1},{\"id\":1,\"title\":\"B\",\"status\":\"done\",\"priority\":2}]}", true},
		{"valid empty", "{\"next_id\":1,\"tasks\":[]}", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "data.json")
			if err := os.WriteFile(path, []byte(tc.content), 0600); err != nil {
				t.Fatal(err)
			}
			_, err := Load(path)
			if (err != nil) != tc.wantErr {
				t.Fatalf("got %v, want error=%t", err, tc.wantErr)
			}
		})
	}
}

func TestMissingFileIsNewStore(t *testing.T) {
	state, err := Load(filepath.Join(t.TempDir(), "absent.json"))
	if err != nil || state.NextID != 1 || len(state.Tasks) != 0 {
		t.Fatalf("got %+v, %v", state, err)
	}
}

func TestInvalidStateDoesNotOverwrite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	if err := os.WriteFile(path, []byte("original"), 0600); err != nil {
		t.Fatal(err)
	}
	err := Save(path, tasks.State{NextID: 0})
	if !errors.Is(err, tasks.ErrInvalid) {
		t.Fatalf("expected ErrInvalid: %v", err)
	}
	content, err := os.ReadFile(path)
	if err != nil || strings.TrimSpace(string(content)) != "original" {
		t.Fatalf("file changed: %q; err=%v", content, err)
	}
}
