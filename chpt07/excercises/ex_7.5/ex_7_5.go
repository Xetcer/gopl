package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type myReader struct {
	r        io.Reader
	n, limit int64
}

func (c *myReader) Read(p []byte) (n int, err error) {
	n, err = c.r.Read(p[:c.limit])
	c.n += int64(n)
	if c.n >= c.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &myReader{r: r, limit: n}
}

func main() {
	lr := LimitReader(strings.NewReader("123456789"), 3)
	b, err := io.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%v", err)
	}
	fmt.Printf("%s\n", b)
}
