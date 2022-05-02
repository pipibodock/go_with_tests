package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Person", "English")
		want := "Hello, Person"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say Hello in Spanish", func(t *testing.T) {
		got := Hello("Person", "Spanish")
		want := "Holla, Person"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say Hello in French", func(t *testing.T) {
		got := Hello("Person", "French")
		want := "Bonjour, Person"
		assertCorrectMessage(t, got, want)
	})
}