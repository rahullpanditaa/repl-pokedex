package tests

import (
	"slices"
	"testing"

	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

func TestCleanInput(t *testing.T) {
	assertSlices := func(t testing.TB, got, want []string, input string) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %s", got, want, input)
		}
	}
	t.Run("uppercase words", func(t *testing.T) {
		inputString := "ThIs teXT has some UPPERCASE leTTERS"
		got := utils.CleanInput(inputString)
		want := []string{"this", "text", "has", "some", "uppercase", "letters"}
		assertSlices(t, got, want, inputString)
	})
	t.Run("sanitized text", func(t *testing.T) {
		inputString := "hello world"
		got := utils.CleanInput(inputString)
		want := []string{"hello", "world"}
		assertSlices(t, got, want, inputString)
	})
	t.Run("uppercase and extra whitespace", func(t *testing.T) {
		inputString := "   This TEXT has leading And trailing whitespace   "
		got := utils.CleanInput(inputString)
		want := []string{"this", "text", "has", "leading", "and", "trailing", "whitespace"}
		assertSlices(t, got, want, inputString)
	})

}
