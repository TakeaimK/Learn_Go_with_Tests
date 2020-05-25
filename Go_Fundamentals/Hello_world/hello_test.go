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
		got := Hello("Sejin", "")
		want := "Hello, Sejin"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spain")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Macron", "France")
		want := "Bonjour, Macron"
		assertCorrectMessage(t, got, want)
	})

}

/*
1. test 작성
2. 컴파일 확인
3. test 실행 및 오류 메세지 확인
4. test 통과를 위해 코드 작성
5. 리팩토링
*/
