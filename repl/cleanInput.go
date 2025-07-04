package repl

import "strings"

func cleanInput(text string) []string {
	textLowercase := strings.Trim(strings.ToLower(text), " ")
	return strings.Split(textLowercase, " ")
}
