package popcount

import "testing"

const bin = 0x1234567890ABCDEF

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(bin))
	}
}

func BenchmarkPopCount_11_6(b *testing.B) {
	bench(b, PopCount)
}

func BenchmarkTableLoop(b *testing.B) {
	bench(b, PopCount_2_3)
}

func BenchmarkPopCountShiftValue(b *testing.B) {
	bench(b, PopCount_2_4)
}

func BenchmarkPopCountDiscardBit(b *testing.B) {
	bench(b, PopCount_2_5)
}
