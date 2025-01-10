// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//go:build ignore
// +build ignore

// Run with:
//
//	$ go run $GOPATH/src/gopl.io/ch8/thumbnail/main.go
//	foo.jpeg
//	^D
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopl/chpt08/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
