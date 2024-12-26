package main

//  go run tempflag.go -temp 273.15K

import (
	"flag"
	"fmt"

	"gopl/chpt07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "температура")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
