package repl

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	inputString := "hello world"
	got := cleanInput(inputString)
	want := []string{"hello", "world"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v given, %s", got, want, inputString)
	}
}
