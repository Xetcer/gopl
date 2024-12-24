package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	words := cmnCounter(p, bufio.ScanWords)
	*c += WordCounter(words)
	return words, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	lines := cmnCounter(p, bufio.ScanLines)
	*c += LineCounter(lines)
	return lines, nil
}

func cmnCounter(p []byte, splitFunc bufio.SplitFunc) int {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(splitFunc)
	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
}

func main() {
	var wordCounter WordCounter
	var byteCounter ByteCounter
	var lineCounter LineCounter
	str := "W1 w2 w3 w4 w5 w6."
	byteCounter.Write([]byte(str))
	wordCounter.Write([]byte(str))
	lineCounter.Write([]byte(str))
	fmt.Println(byteCounter) // 18
	fmt.Println(wordCounter) // 6
	fmt.Println(lineCounter) // 1
}
