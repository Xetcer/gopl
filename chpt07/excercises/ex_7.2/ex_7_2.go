package main

import (
	"fmt"
	"io"
)

type ByteCounter struct {
	w     io.Writer
	count int64
}

// Реализуем интерфейс io.Writer для нашего типа
func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c, &c.count
}

func main() {
	writer, counter := CountingWriter(io.Discard)
	fmt.Fprintf(writer, "test string to count bytes")
	fmt.Println("Bytes count= ", *counter)
}
