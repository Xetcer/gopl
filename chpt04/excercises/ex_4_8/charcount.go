// Charcount вычисляет количество символов Unicode,
// go run charcount.go
// вводишь строки сколько надо
// нажимаешь ctrl+Z чтобы получить символ io.EOF и нажимаешь Enter чтобы программа отработала вывод
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)      // количество символов в Unicode
	var utflen [utf8.UTFMax + 1]int   // Количество длин кодировок UTF-8
	statCount := make(map[string]int) // Количество букв, цифр в строках
	invalid := 0                      // Количество некорректных символов UTF-8

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возврат руны, размера и ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsDigit(r) {
			statCount["Digit"]++
		} else if unicode.IsLetter(r) {
			statCount["Letter"]++
		} else if unicode.IsSpace(r) {
			statCount["Space"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("types\tcount\n")
	for k, value := range statCount {
		fmt.Printf("%q\t%d\n", k, value)
	}
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
}
