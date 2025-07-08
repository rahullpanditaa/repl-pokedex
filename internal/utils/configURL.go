package utils

type Config struct {
	// URLs []struct {
	// 	PreviousURL string
	// 	NextURL     string
	// }
	NextURL     string `json:"next"`
	PreviousURL string `json:"previous"`
}
