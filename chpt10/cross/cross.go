package main

//  go build gopl/chpt10/cross
// ./cross
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
