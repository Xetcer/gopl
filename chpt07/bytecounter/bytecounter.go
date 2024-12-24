package main

import "fmt"

/*
package io
// Writer является интерфейсом, являющимся оболочкой метода Write,
type Writer interface {
// Write записывает len(p) байтов из p в базовый поток данных.
// Метод возвращает количество байтов, записанных из р
// (0 <= п <= 1еп(р)), а любая ошибка вызывает прекращение записи.
// Write должен вернуть ненулевую ошибку при п < 1еп(р).
// Write не должен изменять данные среза, даже временно.
//
// Реализации не должны сохранять р.
Write(p []byte) (n int, err error)
}
*/

type ByteCounter int

// Реализуем интерфейс io.Writer
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Преобразование int в ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5, =len("hello")
	c = 0          // сброс счетчика

	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
