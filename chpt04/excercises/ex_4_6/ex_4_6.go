package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	data := []byte("This is  test  text  ")
	fmt.Printf("Before trim space: %s\n", data)
	fmt.Printf("After trim space: %s\n", convert(data))
}

func convert(data []byte) []byte {
	for i := 0; i < len(data); {
		first, size := utf8.DecodeRune(data[i:]) // получаем первый символ в utf8
		if unicode.IsSpace(first) {              // Если это пробел
			second, _ := utf8.DecodeRune(data[i+size:]) // смотрим второй символ в utf8
			if unicode.IsSpace(second) {                // Если второй тоже пробел
				copy(data[i:], data[i+size:]) // копируем с адреса второго все по адресу первого
				data = data[:len(data)-size]  // обновляем срез с новыми данными
				continue                      // Еще раз проверим с 0 адреса, если ли пробелы после него
			}
		}
		i++
	}
	return data
}
