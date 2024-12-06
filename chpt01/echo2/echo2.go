// вывод аргументов командной строки
// go run echo2.go 1 2 3 - результат 1 2 3
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}
