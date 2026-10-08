package main

import "fmt"

func maxNumber(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func countPositive(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if number > 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Printf("Максимум: %d\n", maxNumber(7, 4))
	fmt.Printf("Положительных: %d\n", countPositive([]int{-2, 1, 0, 5, 10}))
}
