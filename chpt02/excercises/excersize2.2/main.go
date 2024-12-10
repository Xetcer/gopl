package main

import (
	"fmt"
	"gopl/chpt02/excercises/lenconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		lenght, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s = %s", lenconv.Metr(lenght).String(), lenconv.MToCentim(lenconv.Metr(lenght)).String(), lenconv.MToMilim(lenconv.Metr(lenght)).String())
	}
}
