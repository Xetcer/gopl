package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintFast() {
	fmt.Println(strings.Join(os.Args, " "))
}

func PrintSlow() {
	for i, arg := range os.Args {
		fmt.Printf("arg#%d = %s\n", i, arg)
	}
}

func main() {
	PrintFast()
	PrintSlow()
}
