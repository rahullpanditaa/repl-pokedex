package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

func CommandMapBack(config *utils.Config) error {

	var previousPageUrl string
	if config.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		previousPageUrl = config.PreviousURL
	}
	req, err := http.NewRequest("GET", previousPageUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

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
	} else {
		config.PreviousURL = *resp.Previous
	}
	if resp.Next == nil {
		config.NextURL = ""
	} else {
		config.NextURL = *resp.Next
	}

	return nil
}
