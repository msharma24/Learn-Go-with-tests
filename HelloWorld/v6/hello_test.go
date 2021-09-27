package main

import "testing"

func TestHello(t *testing.T) {
  assertCorrectMessage := func(t testing.TB, got , want string) {
    t.Helper()

    if got != want {
      t.Errorf("got %q want %q", got, want)

    }
  }

  t.Run("to a person", func(t *testing.T) {
    got := Hello("Chris", "")
    want := "Hello, Chris"

    assertCorrectMessage(t, got, want)

  })

  t.Run("empty string", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"

    assertCorrectMessage(t, got, want)
  })

  t.Run("saying in spanish", func(t *testing.T) {
    got := Hello("Chris", spanish)
    want := "Hola, Chris"

    assertCorrectMessage(t, got, want)
  })

  t.Run("saying in french", func(t *testing.T) {
    got := Hello("Chris", french)
    want := "Bonjour, Chris"

    assertCorrectMessage(t, got, want)

  })

}
