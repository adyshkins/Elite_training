package main

import "fmt"

func main() {
	var pages int
	fmt.Print("Введите число страниц: ")

	// Чтение проверяем до использования полученного значения.
	_, err := fmt.Scan(&pages)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if pages < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным.")
		return
	}
	fmt.Printf("Будет обработано страниц: %d\n", pages)
}
