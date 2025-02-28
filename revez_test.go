package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzTestRevez(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!123456"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, original string) {
		rev := Revez(original)
		doubleRev := Revez(rev)
		if original != doubleRev {
			t.Errorf("Antes: %q ,Despues: %q", original, doubleRev)
		}
		if utf8.ValidString(original) && !utf8.ValidString(rev) {
			t.Errorf("El reves es invalido %q", rev)
		}
	})
}
