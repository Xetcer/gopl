package main

import "testing"

func TestPrepareComma(t *testing.T) {
	tests := []struct {
		str, want string
	}{
		{"123", "123"},
		{"+123", "+123"},
		{"-123", "-123"},
		{"123456", "123,456"},
		{"+123456", "+123,456"},
		{"-123456", "-123,456"},
		{"123.456", "123.456"},
		{"12356.789", "12,356.789"},
		{"+12356.789", "+12,356.789"},
		{"-12356.789", "-12,356.789"},
		{"123456789", "123,456,789"},
		{"123456789.0", "123,456,789.0"},
		{"+123456789.0", "+123,456,789.0"},
		{"-123456789.0", "-123,456,789.0"},
	}
	for _, test := range tests {
		got := prepareComma(test.str)
		if test.want != got {
			t.Errorf("prepareComma (%q), got %q, want %q", test.str, got, test.want)
		}
	}
}
