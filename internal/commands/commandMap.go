package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

type Response struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(config *utils.Config) error {
	// display the names of 20 location area in the pokemon
	// world

	// create a request
	req, err := http.NewRequest("GET", utils.BaseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// reading response body into a slice of bytes
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var resp Response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resp); err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	if resp.Previous == nil {
		config.PreviousURL = ""
	}
	config.NextURL = *resp.Next

	return nil
}
