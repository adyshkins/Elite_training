package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotFound = errors.New("задача не найдена")
	ErrInvalid  = errors.New("неверные данные задачи")
)

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

type TaskStore interface {
	Add(title string, priority int) (Task, error)
	Get(id int) (Task, error)
	Complete(id int) error
	Delete(id int) error
	List() []Task
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

func readLine(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func report(err error) {
	if err == nil {
		return
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Такой задачи нет")
	} else if errors.Is(err, ErrInvalid) {
		fmt.Println("Неверное название или приоритет")
	} else {
		fmt.Println("Ошибка:", err)
	}
}

func run(store TaskStore, scanner *bufio.Scanner) {
	defer fmt.Println("Работа приложения завершена")
	for {
		fmt.Println("\n1 — добавить | 2 — список | 3 — завершить | 4 — удалить | 0 — выход")
		cmd, ok := readLine(scanner, "> ")
		if !ok || cmd == "0" {
			return
		}
		switch cmd {
		case "1":
			title, ok := readLine(scanner, "Название: ")
			if !ok {
				return
			}
			raw, ok := readLine(scanner, "Приоритет (1–3): ")
			if !ok {
				return
			}
			priority, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("Приоритет должен быть целым числом")
				continue
			}
			task, err := store.Add(title, priority)
			if err != nil {
				report(err)
				continue
			}
			fmt.Printf("Создана задача #%d\n", task.ID)
		case "2":
			tasks := store.List()
			if len(tasks) == 0 {
				fmt.Println("Список пуст")
			}
			for _, task := range tasks {
				fmt.Printf("#%d %s [%s], приоритет %d\n", task.ID, task.Title, task.Status, task.Priority)
			}
		case "3", "4":
			raw, ok := readLine(scanner, "ID: ")
			if !ok {
				return
			}
			id, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("ID должен быть целым числом")
				continue
			}
			if cmd == "3" {
				err = store.Complete(id)
			} else {
				err = store.Delete(id)
			}
			if err != nil {
				report(err)
			} else {
				fmt.Println("Готово")
			}
		default:
			fmt.Println("Неизвестная команда. Выберите 0–4")
		}
	}
}

func main() {
	run(NewMemoryStore(), bufio.NewScanner(os.Stdin))
}
