package main

import "fmt"

/*
Вариативные функции
*/

/*
	sum - вариативная функция, принимающая в качестве аргументов значения типа int, любое количество
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
	values := []int{1, 2, 3, 4}
	// Если передаем срез, то вызываем функцию со срезом values...
	fmt.Println(sum(values...)) // "10"
}
