// вывод аргументов командной строки
// go run echo3.go 1 2 3 - результат 1 2 3
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
