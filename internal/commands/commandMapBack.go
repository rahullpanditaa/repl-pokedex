package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rahullpanditaa/repl-pokedex/internal/pokecache"
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

	// cache has url and corresponding list of 20 location areas
	apiResp, respExists := pokecache.ApiCache.Get(previousPageUrl)
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

		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&resp); err != nil {
			log.Fatal(err)
		}

		var locationAreas []string
		for _, area := range resp.Results {
			fmt.Println(area.Name)
			locationAreas = append(locationAreas, area.Name)
		}
		// pokecache.ApiCache.Add(previousPageUrl, []byte(locationAreas))
		cacheData, err := json.Marshal(locationAreas)
		if err != nil {
			log.Fatal(err)
		}
		pokecache.ApiCache.Add(previousPageUrl, cacheData)

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
