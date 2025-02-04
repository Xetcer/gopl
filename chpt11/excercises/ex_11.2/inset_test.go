package inset

import (
	"fmt"
	"testing"
)

const (
	AddTest   = 0
	HasTest   = 1
	UnionTest = 2
)

func TestInset(t *testing.T) {
	var tests = []struct {
		testName string
		args     []int
		testType int
		want     string
	}{
		{testName: "Add data", args: []int{1, 144, 9}, testType: AddTest, want: "{1 9 144}"},
		{testName: "Has data", args: []int{1, 144, 9}, testType: HasTest, want: "true"},
		{testName: "Union with", args: []int{1, 144, 9}, testType: UnionTest, want: "{1 9 42 144}"},
	}
	for _, test := range tests {
		fmt.Println(test.testName)
		switch test.testType {
		case AddTest:
			{
				t.Run(test.testName, func(t *testing.T) {
					var x IntSet
					for _, arg := range test.args {
						x.Add(arg)
					}
					if x.String() != test.want {
						t.Errorf("Add(x)=%s, but %s expected", x.String(), test.want)
					}
				})
			}
		case HasTest:
			{
				t.Run(test.testName, func(t *testing.T) {
					var x IntSet
					for _, arg := range test.args {
						x.Add(arg)
					}
					if fmt.Sprint(x.Has(9)) != test.want {
						t.Errorf("Has(x)=%s, but %s expected", fmt.Sprint(x.Has(9)), test.want)
					}
				})
			}
		case UnionTest:
			{
				t.Run(test.testName, func(t *testing.T) {
					var x, y IntSet
					for _, arg := range test.args {
						x.Add(arg)
					}
					y.Add(9)
					y.Add(42)
					x.UnionWith(&y)
					got := x.String()
					if got != test.want {
						t.Errorf("Has(x)=%s, but %s expected", got, test.want)
					}
				})
			}
		}
	}
}
