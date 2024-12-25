package main

/*
go run outline2.go http://gopl.io
*/
import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type UrlReader struct {
	Url string
	i   int64
}

func (c *UrlReader) Read(b []byte) (n int, err error) {
	if c.i >= int64(len(c.Url)) {
		return 0, io.EOF
	}
	n = copy(b, c.Url[c.i:])
	c.i = int64(n)
	fmt.Printf("Parsed url: %s\n", c.Url)
	return
}

func NewReader(s string) io.Reader {
	return &UrlReader{s, 0}
}

func main() {
	doc, err := html.Parse(NewReader("sdelanounas.ru"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Внесение дескриптора в стек
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
