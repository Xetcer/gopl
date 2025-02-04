// Упражнение 11.5. Расширьте TestSplit так, чтобы она использовала таблицу входных и ожидаемых выходных данных.
package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		testName string
		args     struct{ s, sep string }
		want     int
	}{
		{testName: "One sign", args: struct {
			s   string
			sep string
		}{s: "a", sep: ":"}, want: 1},
		{testName: "Three sign", args: struct {
			s   string
			sep string
		}{s: "a:b:c", sep: ":"}, want: 3},
		{testName: "Zero sign", args: struct {
			s   string
			sep string
		}{s: "", sep: ""}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			words := strings.Split(test.args.s, test.args.sep)
			if got := len(words); got != test.want {
				t.Errorf("Split(%q, %q) возвращает %d слов, а требуется %d", test.args.s, test.args.sep, got, test.want)
			}
		})
	}

}
