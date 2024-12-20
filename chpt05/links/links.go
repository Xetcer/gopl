package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url) // Отправляем запрос
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Получение %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body) // Получаем узлы HTML разметки
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("анализ %s как HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" { // Если это элемент разметки и это ТЭГ ссылка <a>
			for _, a := range n.Attr { // читаем все атрибуты ссылки
				if a.Key != "href" { // Если находим атрибут не находим href
					continue // переходим к следующему атрибуту
				}
				link, err := resp.Request.URL.Parse(a.Val) // Пытаемся распарсить адрес ссылки
				if err != nil {                            // Если не получилось
					continue // продолжаем
				}
				links = append(links, link.String()) // добавляем в список ссылок новую найденную
			}
		}
	}
	forEachNode(doc, visitNode, nil) // Перебираем все узлы, передавая в качестве pre функции visitNode
	return links, nil
}

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
