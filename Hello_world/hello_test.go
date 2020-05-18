package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	//에러 발생 시 출력되는 공유 코드
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say my name", func(t *testing.T) {
		got := Hello("Sejin")
		want := "Hello, Sejin"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

}
