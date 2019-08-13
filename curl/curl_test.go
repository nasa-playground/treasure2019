package main

import (
	"testing"
)

func TestUrlParse(t *testing.T) {
	var flagtests = []struct {
		in  string
		out bool
	}{
		{"", false},
		{"http://?name=hogehoge", false},

		{"https://google.com", true},
		{"http://google.com", true},
		{"http://google.com/?name=hogehoge", true},
	}

	for _, tt := range flagtests {
		if ParseUrl(tt.in) != tt.out {
			t.Fatal(tt.in, ParseUrl(tt.in))
		}
	}
}
