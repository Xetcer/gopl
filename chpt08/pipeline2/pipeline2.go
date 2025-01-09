/*
		пример Конвейера
		Первая программа генерирует целые числа 0, 1,2, ... и отправляет их по каналу второй go-подпрограмме,
		которая получает значения, возводит их в квадрат и передает по следующему каналу третьей go-подпрограмме,
	 	которая выводит получаемые значения на экран.
*/
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Генерация
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//вывод (в главной подпрограмме)
	for x := range squares {
		fmt.Println(x)
	}
}
