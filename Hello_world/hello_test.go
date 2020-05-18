package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say my name", func(t *testing.T) {
		got := Hello("Sejin")
		want := "Hello, Sejin"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Say 'Hello world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
