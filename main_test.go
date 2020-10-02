package main

import "testing"

func TestHello(t *testing.T) {
	observed := Hello()
	expected := "Hello, Shreyas"

	if observed != expected {
		t.Errof("got %q, want %q", observed, expected)
	}
}