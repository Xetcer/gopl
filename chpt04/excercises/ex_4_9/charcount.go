// Charcount вычисляет количество символов Unicode,
// go run charcount.go
// вводишь строки сколько надо
// нажимаешь ctrl+Z чтобы получить символ io.EOF и нажимаешь Enter чтобы программа отработала вывод
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // тут будем хранить все слова, которые будут введены в виде ключей и считать их число

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Printf("%q\t\t%q\n", "word", "count")
	for key, value := range counts {
		fmt.Printf("%q\t\t%d\n", key, value)
	}
}
