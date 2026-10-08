package main

import "fmt"

func main() {
	statuses := []string{"new", "done", "in_progress"}

	fmt.Println("Обычный for:")
	for i := 0; i < len(statuses); i++ {
		fmt.Printf("%d: %s\n", i+1, statuses[i])
	}

	fmt.Println("for range:")
	for i, status := range statuses {
		fmt.Printf("%d: %s\n", i+1, status)
	}

	remaining := 3
	for remaining > 0 {
		remaining--
	}
	fmt.Printf("Осталось: %d\n", remaining)
}
