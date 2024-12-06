// Dup2 выводит текст каждой строки, которая появляется во
// входных данных более одного раза. Программа читает
// стандартный ввод или список именованных файлов
// go run dup2.go - вводим строки, жмем ctrl+z В конце.
//
// Будут выведены повторные строки и число повторов.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			lineConuts := len(counts)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			if lineConuts < len(counts) {
				fmt.Printf("In file %s found repeated lines!\n", arg)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// игнорируем потенциальные ошибки input.Err()
}
