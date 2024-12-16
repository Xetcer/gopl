package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	size := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[size-i] = s[size-i], s[i]
	}
}

func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:]) // определяем текущий символ в кодировке UTF-8
		// fmt.Printf("%s len is %d\n", b[i:i+size], size)
		reverse(b[i : i+size]) // Разворачиваем текущий символ в кодировке
		i += size              // Переходим к следующему символу кодировки
	}
	reverse(b) // Развернем еще раз весь срез, чтобы получить корректную последовательность байт в развернутом срезе
	return b   // Вернем срез
}

func main() {
	a := "一 二 三"
	fmt.Printf("%s\n", a)

	fmt.Printf("%s\n", revUTF8([]byte(a)))
}
