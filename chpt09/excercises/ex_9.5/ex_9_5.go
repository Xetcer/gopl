package main

import (
	"fmt"
	"sync"
	"time"
)

// канал для широковещательного закрытия подпрограмм
var chClose = make(chan struct{})
var wg = sync.WaitGroup{}

func pingPong(id int, in, out chan string) {
	defer wg.Done()
	for {
		select {
		case <-chClose:
			return
		case msgIn, ok := <-in:
			if ok {
				fmt.Println(id, " goroutine recv msg: ", msgIn)
				switch id {
				case 1:
					out <- "ping"
				case 2:
					out <- "pong"
				}
			} else {
				return
			}
		}
	}
}

func main() {
	// for debug
	// ch1 := make(chan string)
	// defer close(ch1)
	// ch2 := make(chan string)
	// defer close(ch2)
	// wg.Add(2)
	// go pingPong(1, ch1, ch2)
	// go pingPong(2, ch2, ch1)

	// ch1 <- "ping"
	// <-time.After(1 * time.Second)
	// close(chClose)
	// wg.Wait()
	// fmt.Println("Program is finished!")

	pings := make(chan string)
	pongs := make(chan string)
	var i int

	start := time.Now()
	go func() {
		for {
			i++
			pings <- "ping"
			<-pongs
		}
	}()

	go func() {
		for {
			i++
			<-pings
			pongs <- "pong"
		}
	}()

	<-time.After(10 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("%.f op/s. i=%d, t=%s", float64(i)/elapsed.Seconds(), i, elapsed)

}
