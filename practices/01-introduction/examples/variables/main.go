package main

import "fmt"

func main() {
	var language string = "Go"
	completed := 2
	const practiceTitle = "Первая программа"
	var ready bool
	var title string

	completed = 3
	fmt.Printf("Язык: %s\n", language)
	fmt.Printf("Практика: %s\n", practiceTitle)
	fmt.Printf("Завершено шагов: %d\n", completed)
	fmt.Printf("Нулевые значения: ready=%t, title=%q\n", ready, title)

	a := 17
	b := 5
	fmt.Printf("Целочисленное деление: %d\n", a/b)
	fmt.Printf("Остаток: %d\n", a%b)
	fmt.Printf("Дробное деление: %.2f\n", float64(a)/float64(b))
}
