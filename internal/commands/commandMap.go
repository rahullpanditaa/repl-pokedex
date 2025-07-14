package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rahullpanditaa/repl-pokedex/internal/pokecache"
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

func CommandMapForward(config *utils.Config) error {
	var nextPageUrl string
	if config.NextURL == "" {
		nextPageUrl = utils.BaseURL
	} else {
		nextPageUrl = config.NextURL
	}

	apiResp, respExists := pokecache.ApiCache.Get(nextPageUrl)
	if respExists {
		fmt.Println("Retrieving from cache...")
		var locationAreas []string
		err := json.Unmarshal(apiResp, &locationAreas)
		if err != nil {
			log.Fatal(err)
		}
		for _, area := range locationAreas {
			fmt.Println(area)
		}
		return nil
	} else {
		var resp Response
		req, err := http.NewRequest("GET", nextPageUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&resp); err != nil {
			log.Fatal(err)
		}

		var locationAreas []string
		for _, area := range resp.Results {
			fmt.Println(area.Name)
			locationAreas = append(locationAreas, area.Name)
		}

		cacheData, err := json.Marshal(locationAreas)
		if err != nil {
			log.Fatal(err)
		}
		pokecache.ApiCache.Add(nextPageUrl, cacheData)

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
	}
	return nil
}
