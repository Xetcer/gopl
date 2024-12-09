package main

import "fmt"

func main() {
	const freezigF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezigF, fToC(freezigF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
