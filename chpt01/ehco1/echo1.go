// Echol выводит аргументы командной строки
// go run echo1.go 1 2 3  вывод 1 2 3
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
