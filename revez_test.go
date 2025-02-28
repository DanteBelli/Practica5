package main

import (
	"testing"
)

func TestRevez(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!123456", "654321!"},
	}
	for _, tc := range testcases {
		rev := Revez(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q , want %q", rev, tc.want)
		}
	}
}
