package utils

import "strings"

func CleanInput(text string) []string {
	textLowercase := strings.Trim(strings.ToLower(text), " ")
	return strings.Split(textLowercase, " ")
}
