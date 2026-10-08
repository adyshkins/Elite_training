package main

import "fmt"

func main() {
	status := "in_progress"
	attempts := 2

	if attempts > 0 {
		fmt.Println("Попытки остались: да")
	} else {
		fmt.Println("Попытки остались: нет")
	}

	switch status {
	case "new":
		fmt.Println("Статус: новая")
	case "in_progress":
		fmt.Println("Статус: в работе")
	case "done":
		fmt.Println("Статус: завершена")
	default:
		fmt.Println("Статус: неизвестен")
	}
}
