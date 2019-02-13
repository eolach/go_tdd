package main

import "testing"

func TestHello(t *testing.T) {

	// Requirement is that the application should greet the user
	// with their name if provided and in aother language if specified

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")          // First run the function
		want := "Hello, Christ"            // Then state the expected result
		assertCorrectMessage(t, got, want) //and test the result
	})

	t.Run("Say 'Hello, world, if name is blank", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("is Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
