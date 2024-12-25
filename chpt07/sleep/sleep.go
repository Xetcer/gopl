package main

//  go run sleep.go --period="10s"
import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "Sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Ожидание %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
