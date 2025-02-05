package inset

// go test -bench=Benchmark

import (
	"math/rand"
	"testing"
	"time"
)

var (
	s1 []int
	s2 []int
)

const (
	n     = 100000
	scale = 100
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.New(rand.NewSource(seed))
	s1 = randInt(n)
	s2 = randInt(n)
}

func randInt(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(scale * n)
	}
	return ints
}

func BenchmarkIntSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := &IntSet{}
		for _, v := range s1 {
			s.Add(v)
		}
	}
}

func BenchmarkIntSetHas(b *testing.B) {
	s := &IntSet{}
	for _, v := range s1 {
		s.Add(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range s1 {
			s.Has(v)
		}
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	is1 := &IntSet{}
	for _, v := range s1 {
		is1.Add(v)
	}
	is2 := &IntSet{}
	for _, v := range s2 {
		is2.Add(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is1.UnionWith(is2)
	}
}

func BenchmarkMapAdd(b *testing.B) {
	s := make(map[int]bool)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range s1 {
			s[k] = true
		}
	}
}

func BenchmarkMapHas(b *testing.B) {
	s := make(map[int]bool)
	for _, k := range s1 {
		s[k] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range s1 {
			_ = s[k]
		}
	}
}

func BenchmarkMapUnionWith(b *testing.B) {
	ms1 := make(map[int]bool)
	for _, k := range s1 {
		ms1[k] = true
	}
	ms2 := make(map[int]bool)
	for _, k := range s2 {
		ms2[k] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range ms2 {
			ms1[k] = true
		}
	}
}
