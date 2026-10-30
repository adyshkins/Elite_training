package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"example.com/go-practice-01/internal/persistence"
	"example.com/go-practice-01/internal/tasks"
)

func ask(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return strings.TrimSpace(scanner.Text()), nil
}

func report(err error) {
	switch {
	case errors.Is(err, tasks.ErrNotFound):
		fmt.Println("Задача не найдена")
	case errors.Is(err, tasks.ErrInvalid):
		fmt.Println("Некорректные данные задачи")
	case err != nil:
		fmt.Println("Ошибка:", err)
	}
}

func run() error {
	path := os.Getenv("TASKS_FILE")
	if path == "" {
		path = "tasks.json"
	}
	state, err := persistence.Load(path)
	if err != nil {
		return err // Do not overwrite corrupted data.
	}
	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	defer fmt.Println("Приложение завершено")
	fmt.Println("Файл данных:", path)
	for {
		fmt.Println("\n1 — добавить; 2 — список; 3 — завершить; 4 — удалить; 0 — выход")
		command, err := ask(scanner, "Команда: ")
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		switch command {
		case "0":
			return nil
		case "1":
			title, err := ask(scanner, "Название: ")
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return err
			}
			raw, err := ask(scanner, "Приоритет (1–3): ")
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return err
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
			if err := persistence.Save(path, store.Snapshot()); err != nil {
				return fmt.Errorf("сохранение после добавления: %w", err)
			}
			fmt.Println("Добавлена задача ID:", task.ID)
		case "2":
			list := store.List()
			if len(list) == 0 {
				fmt.Println("Список пуст")
			}
			for _, t := range list {
				fmt.Printf("#%d | %s | %s | приоритет %d\n", t.ID, t.Title, t.Status, t.Priority)
			}
		case "3", "4":
			raw, err := ask(scanner, "ID: ")
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return err
			}
			id, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("ID должен быть целым числом")
				continue
			}
			if command == "3" {
				err = store.Complete(id)
			} else {
				err = store.Delete(id)
			}
			if err != nil {
				report(err)
				continue
			}
			if err := persistence.Save(path, store.Snapshot()); err != nil {
				return fmt.Errorf("сохранение после изменения: %w", err)
			}
			fmt.Println("Изменение сохранено")
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка приложения:", err)
		os.Exit(1)
	}
}
