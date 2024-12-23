package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//Fetch загружает URL и возвращает имя и длину локального файла.
//  go run fetch.go http://www.gopl.io

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Закрытие файла, если ошибка Copy , возвращаем ее
	// Если ошибка есть в f.Close и нет в Copy - возвращаем f.Close
	// Если ошибка есть и в f.Close и в Copy - возвращаем Copy
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
