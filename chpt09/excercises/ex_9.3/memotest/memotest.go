// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

// !+httpRequestBody
func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

type MCancel interface {
	Get(key string, done chan struct{}) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m MCancel) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

func Concurrent(t *testing.T, m MCancel) {
	//!+conc
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}

func ConcurrentWithCancelDelay(t *testing.T, m MCancel) {
	var n sync.WaitGroup
	cancelCh := make(chan struct{})
	var firstMu sync.Mutex
	isFirst := true
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			cancel := make(chan struct{})
			firstMu.Lock()
			if isFirst {
				cancel = cancelCh
				isFirst = false
			}
			firstMu.Unlock()
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			if value != nil {
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
			} else {
				fmt.Printf("%s, %s, operation is canceled\n",
					url, time.Since(start))
			}
		}(url)
	}
	// time.Sleep(1 * time.Second)
	close(cancelCh)
	n.Wait()
}
