package popcount

// go test -bench=BenchmarkPopCount выводит таблицу выполнения всех функций BenchmarkPopCount
import "testing"

var N = 1000

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < N; i++ {
		PopCount(100)
	}
}

func BenchmarkPopCount_2_3(b *testing.B) {
	for i := 0; i < N; i++ {
		PopCount_2_3(100)
	}
}

func BenchmarkPopCount_2_4(b *testing.B) {
	for i := 0; i < N; i++ {
		PopCount_2_4(100)
	}
}

func BenchmarkPopCount_2_5(b *testing.B) {
	for i := 0; i < N; i++ {
		PopCount_2_5(100)
	}
}
