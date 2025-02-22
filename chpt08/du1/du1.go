package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
walkDir рекурсивно обходит дерево файлов с корнем в dir
и отправляет размер каждого найденного файла в fileSizes.
*/
func walkdir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkdir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	return entries
}

func main() {
	// определяет исходные каталоги
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// обход дерева файлов
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkdir(root, fileSizes)
		}
		close(fileSizes)
	}()
	// вывод результатов
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d файлов %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
