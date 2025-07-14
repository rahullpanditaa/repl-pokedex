package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rahullpanditaa/repl-pokedex/internal/pokecache"
	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

type AreaResponse []struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int   `json:"chance"`
			ConditionValues []any `json:"condition_values"`
			MaxLevel        int   `json:"max_level"`
			Method          struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"method"`
			MinLevel int `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int `json:"max_chance"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
}

func CommandExplore(config *utils.Config, areaName string) error {
	areaUrl := utils.BaseURL + areaName
	areaDetails, areaExists := pokecache.ApiCache.Get(areaUrl)
	if areaExists {
		fmt.Println("Retrieving from cache...")
		var pokemonsInArea []string
		err := json.Unmarshal(areaDetails, &pokemonsInArea)
		if err != nil {
			log.Fatal(err)
		}
		for _, pokemon := range pokemonsInArea {
			fmt.Println(pokemon)
		}
		return nil
	} else {
		var resp AreaResponse
		req, err := http.NewRequest("GET", areaUrl, nil)
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
		for _, area := range resp {
			fmt.Println("Exploring " + areaName + "...")
			fmt.Println("Found pokemon:")
			fmt.Println(" - " + area.Pokemon.Name)
			locationAreas = append(locationAreas, area.Pokemon.Name)
		}
		cacheData, err := json.Marshal(locationAreas)
		if err != nil {
			log.Fatal(err)
		}
		pokecache.ApiCache.Add(areaUrl, cacheData)
	}
	return nil
}
