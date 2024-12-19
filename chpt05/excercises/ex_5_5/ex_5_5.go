/*
go run ex_5_5.go https://golang.org
подсчитываем сколько слов и картинок во HTMl странице по указанному адресу.
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func countWordsAndImages(url string) (words, images int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	words, images = countWordsImages(doc)
	return
}

func countWordsImages(n *html.Node) (words, images int) {
	switch n.Type {
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) > 0 && n.Parent.Data != "script" && n.Parent.Data != "style" {
			words += len(strings.Fields(n.Data))
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsImages(c)
		words += w
		images += i
	}
	return words, images
}

func main() {
	url := os.Args[1]
	words, images := countWordsAndImages(url)
	fmt.Printf("words=%d, images=%d", words, images)
}
