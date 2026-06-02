package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Henry")
	want := "Hello, Henry"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
