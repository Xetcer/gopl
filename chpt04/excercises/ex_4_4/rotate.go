/*
Exercise 4.3
Напишите версию функции rotate, которая работает в один проход.
*/

package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(rotate(a[:], 1))
}

// rotate делает сдвиг элементов на указанное количество позиций
func rotate(slice []int, i int) []int {
	i %= len(slice)                    // получаем крайний элемент заданного сдвига
	tmp := append(slice, slice[:i]...) // добавляем в переменную tmp диапазон значений до крайнего элемента
	copy(slice, tmp[i:])               // копированием отбрасываем значения исходного среза.
	return slice

}
