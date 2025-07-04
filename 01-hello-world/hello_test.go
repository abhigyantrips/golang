package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to myself", func(t *testing.T) {
		got := Hello("Abhigyan", "")
		want := "Hello, Abhigyan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Abhigyan", "Spanish")
		want := "Hola, Abhigyan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Abhigyan", "French")
		want := "Bonjour, Abhigyan"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}