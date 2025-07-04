package repl

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	t.Run("uppercase words", func(t *testing.T) {
		inputString := "ThIs teXT has some UPPERCASE leTTERS"
		got := cleanInput(inputString)
		want := []string{"this", "text", "has", "some", "uppercase", "letters"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %s", got, want, inputString)
		}
	})
	t.Run("sanitized text", func(t *testing.T) {
		inputString := "hello world"
		got := cleanInput(inputString)
		want := []string{"hello", "world"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %s", got, want, inputString)
		}
	})
	t.Run("uppercase and extra whitespace", func(t *testing.T) {
		inputString := "   This TEXT has leading And trailing whitespace   "
		got := cleanInput(inputString)
		want := []string{"this", "text", "has", "leading", "and", "trailing", "whitespace"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %s", got, want, inputString)
		}
	})

}
