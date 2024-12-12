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
)

func basename(s string) string {
	// отбрасываем последний символ '/' и все перед ним.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// сохраняем все до последней точки
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}
