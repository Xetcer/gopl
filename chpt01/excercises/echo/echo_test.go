package main

import "testing"
var N = 1000
func BenchmarkPrintFast(b *testing.B) {
	for i := 0; i < N; i++ {
		PrintFast()
	}
}

func BenchmarkPrintSlow(b *testing.B) {
	for i := 0; i < N; i++ {
		PrintSlow()
	}
}
