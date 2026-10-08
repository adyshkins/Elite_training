package main

import "fmt"

func main() {
	fixed := [3]string{"new", "in_progress", "done"}
	statuses := []string{"new", "done", "new"}
	statuses = append(statuses, "done")

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	missing, exists := counts["blocked"]
	fmt.Printf("Массив: %d элемента\n", len(fixed))
	fmt.Printf("Срез: %d элемента\n", len(statuses))
	fmt.Printf("new: %d, done: %d\n", counts["new"], counts["done"])
	fmt.Printf("blocked: значение=%d, ключ существует=%t\n", missing, exists)
}
