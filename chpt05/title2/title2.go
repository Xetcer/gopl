package main

/*
Инструкция defer часто используется с такими парными операциями, как открытие и закрытие,
подключение и отключение или блокировка и разблокирование — для гарантии освобождения
ресурсов во всех случаях, независимо от того, насколько сложен поток управления.
Правильное место инструкции defer, которая освобождает ресурс, — сразу же после того,
как ресурс был успешно захвачен. В функции title ниже один отложенный вызов заменяет
оба предыдущих вызова resp.Body.

go run title2.go http://gopl.io - OK
go run title2.go https://golang.org/doc/gopher/frontpage.png - BAD
*/

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s имеет тип %s, не text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("разбор %s как HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)

	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}
