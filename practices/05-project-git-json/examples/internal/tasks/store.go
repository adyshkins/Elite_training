package tasks

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound = errors.New("задача не найдена")
	ErrInvalid  = errors.New("неверные данные задачи")
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority int    `json:"priority"`
}

// State contains both the tasks and the ID counter.
// Keeping NextID prevents reusing IDs after deletion and restart.
type State struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

type MemoryStore struct {
	tasks  []Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}

func (s *MemoryStore) Add(title string, priority int) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" || priority < 1 || priority > 3 {
		return Task{}, fmt.Errorf("название или приоритет: %w", ErrInvalid)
	}
	task := Task{ID: s.nextID, Title: title, Status: "new", Priority: priority}
	s.tasks = append(s.tasks, task)
	s.nextID++
	return task, nil
}

func (s *MemoryStore) Get(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("id %d: %w", id, ErrNotFound)
}

func (s *MemoryStore) Complete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Status = "done"
			return nil
		}
	}
	return fmt.Errorf("id %d: %w", id, ErrNotFound)
}

func (s *MemoryStore) Delete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("id %d: %w", id, ErrNotFound)
}

func (s *MemoryStore) List() []Task {
	return append([]Task(nil), s.tasks...)
}

func (s *MemoryStore) Snapshot() State {
	return State{NextID: s.nextID, Tasks: s.List()}
}

// Restore checks persisted data BEFORE modifying the store.
func (s *MemoryStore) Restore(state State) error {
	if state.NextID < 1 {
		return fmt.Errorf("next_id должен быть положительным: %w", ErrInvalid)
	}
	seen := make(map[int]bool, len(state.Tasks))
	for _, t := range state.Tasks {
		if t.ID < 1 || t.ID >= state.NextID ||
			strings.TrimSpace(t.Title) == "" || t.Title != strings.TrimSpace(t.Title) ||
			(t.Status != "new" && t.Status != "done") ||
			t.Priority < 1 || t.Priority > 3 || seen[t.ID] {
			return fmt.Errorf("некорректная задача id=%d: %w", t.ID, ErrInvalid)
		}
		seen[t.ID] = true
	}
	s.tasks = append([]Task(nil), state.Tasks...)
	s.nextID = state.NextID
	return nil
}
