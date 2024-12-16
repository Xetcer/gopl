package reverse

import "testing"

func TestReverse(t *testing.T) {

	tests := []struct {
		arr, want [5]int{}
	}{
		{arr: [...]int{1, 2, 3, 4, 5}, want: [...]int{5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		reverse(&test.arr)
		if test.arr != test.want {
			t.Errorf("After rotate %q, but %q expected", test.arr, test.want)
		}
	}
}
