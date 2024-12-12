/*
go run basename1.go
a/b/c.go - c
c.d.go - c.d
abc - abc
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")              // -1 если не найден
	s = s[slash+1:]                                 // отбрасываем все что слева от крайнего правого слеша вместе с ним
	if dot := strings.LastIndex(s, "."); dot >= 0 { // находим крайнюю правую точку и убираем все что за ней вправо вместе с самой точкой
		s = s[:dot]
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}
