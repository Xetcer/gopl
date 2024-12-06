// Dupl выводит текст каждой строки, которая появляется в
// стандартном вводе более одного раза, а также количество
// ее появлений.
// go run dup1.go - вводим строки, если будут одинаковые строки то программа
// выведет количество повторов, при этом для окончания ввода ctrl+z нажать.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Примечание: игнорируем потенциальные
	// ошибки из input.ErrQ
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
