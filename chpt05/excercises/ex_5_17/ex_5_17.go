package main

/*
go run ex_5_17.go
*/

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
// 	elements := make([]*html.Node, 0)
// 	if len(name) == 0 {
// 		return elements
// 	}

// 	var getNodes func(node *html.Node, name string) []*html.Node

// 	getNodes = func(node *html.Node, name string) []*html.Node {
// 		nodes := make([]*html.Node, 0)
// 		if node.DataAtom.String() == name {
// 			nodes = append(nodes, node)
// 		}
// 		if node.NextSibling != nil {
// 			nodes = append(nodes, getNodes(node.NextSibling, name)...)
// 		}
// 		return nodes
// 	}

//		for _, n := range name {
//			elements = append(elements, getNodes(doc, n)...)
//		}
//		return elements
//	}
func ElementsByTagName(doc *html.Node, tags ...string) (nodes []*html.Node) {
	if len(tags) == 0 {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, tag := range tags {
			if doc.Data == tag {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, tags...)...)
	}
	return nodes
}

func main() {
	resp, err := http.Get("https://www.scrapethissite.com/pages/frames/?frame=i")
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(doc)
	nodes := ElementsByTagName(doc, "img", "h3", "a")
	for _, n := range nodes {
		fmt.Println(n.Data)
		fmt.Println(n.Attr)
	}
}
