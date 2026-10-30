package tasks

import (
	"errors"
	"testing"
)

func TestMemoryStoreCRUD(t *testing.T) {
	s := NewMemoryStore()
	if _, err := s.Add("  ", 2); !errors.Is(err, ErrInvalid) {
		t.Fatalf("empty title: %v", err)
	}
	if _, err := s.Add("Valid", 0); !errors.Is(err, ErrInvalid) {
		t.Fatalf("invalid priority: %v", err)
	}
	a, err := s.Add("  Изучить Go  ", 2)
	if err != nil || a.ID != 1 || a.Title != "Изучить Go" {
		t.Fatalf("first add: %+v, %v", a, err)
	}
	if err := s.Complete(a.ID); err != nil {
		t.Fatal(err)
	}
	if err := s.Complete(a.ID); err != nil {
		t.Fatal(err)
	}
	list := s.List()
	if len(list) != 1 || list[0].Status != "done" {
		t.Fatalf("complete: %+v", list)
	}
	list[0].Title = "Изменено снаружи"
	got, _ := s.Get(a.ID)
	if got.Title != "Изучить Go" {
		t.Fatalf("List leaked store state: %+v", got)
	}
	if err := s.Delete(a.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Get(a.ID); !errors.Is(err, ErrNotFound) {
		t.Fatalf("get deleted ID: %v", err)
	}
	if err := s.Delete(a.ID); !errors.Is(err, ErrNotFound) {
		t.Fatalf("delete twice: %v", err)
	}
	b, err := s.Add("Новая задача", 3)
	if err != nil || b.ID != 2 {
		t.Fatalf("ID reused: %+v %v", b, err)
	}
}

func TestRestoreRejectsInvalidStateWithoutChangingStore(t *testing.T) {
	s := NewMemoryStore()
	_, _ = s.Add("Исходная", 1)
	before := s.Snapshot()
	bad := []State{
		{NextID: 0},
		{NextID: 1, Tasks: []Task{{ID: 1, Title: "A", Status: "new", Priority: 1}}},
		{NextID: 3, Tasks: []Task{{ID: 1, Title: "A", Status: "wrong", Priority: 1}}},
		{NextID: 3, Tasks: []Task{{ID: 1, Title: "A", Status: "new", Priority: 1}, {ID: 1, Title: "B", Status: "done", Priority: 2}}},
	}
	for i, state := range bad {
		if err := s.Restore(state); !errors.Is(err, ErrInvalid) {
			t.Fatalf("case %d: %v", i, err)
		}
		after := s.Snapshot()
		if after.NextID != before.NextID || len(after.Tasks) != len(before.Tasks) || after.Tasks[0] != before.Tasks[0] {
			t.Fatalf("case %d modified live state: %+v", i, after)
		}
	}
}
