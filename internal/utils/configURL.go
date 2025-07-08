package utils

type Config struct {
	NextURL     string `json:"next"`
	PreviousURL string `json:"previous"`
}

var Urls Config
