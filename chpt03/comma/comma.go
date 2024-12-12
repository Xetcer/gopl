/*
Задача заключается в том, чтобы взять строковое представление целого числа, такое как ”12345”,
и вставить запятые, разделяющие каждые три цифры, как в строке "12,345".
Эта версия работает только с целыми числами;
go run comma.go 123456 - 123,456
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}
