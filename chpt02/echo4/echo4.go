// Echo4 выводит аргументы командной строки
// использование go run echo4.go -s / a bc def вывод: a/bc/def
// go run echo4.go -help
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
