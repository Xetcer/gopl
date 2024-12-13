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
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func prepareComma(s string) string {
	if len(s) <= 3 {
		return s
	}
	startIndex := 0
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		startIndex = 1
		// fmt.Println("start index is", startIndex)
	}
	pointIndex := strings.Index(s, ".")
	if pointIndex != -1 {
		// fmt.Println("from sign to comma str is:", s[startIndex:pointIndex])
		// fmt.Println("from comma to end str is:", s[pointIndex:])
		s = s[:startIndex] + comma(s[startIndex:pointIndex]) + s[pointIndex:]
	} else {
		s = s[:startIndex] + comma(s[startIndex:])
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(prepareComma(input.Text()))
	}
}
