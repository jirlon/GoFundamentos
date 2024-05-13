package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("dizendo oi para as pessoas", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("string vazia tem como padrão 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	//t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}