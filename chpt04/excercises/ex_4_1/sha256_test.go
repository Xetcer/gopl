package main

import "testing"

func TestCompareBits(t *testing.T) {
	tests := []struct {
		b1, b2 byte
		want   int
	}{
		{b1: 1, b2: 2, want: 2},
		{b1: 1, b2: 3, want: 1},
	}
	for _, test := range tests {
		got := compareBits(test.b1, test.b2)
		if got != test.want {
			t.Errorf("Got %d, want %d", got, test.want)
		}
	}
}
