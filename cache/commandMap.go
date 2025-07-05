package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func CommandMap() error {
	// display the names of 20 location area in the pokemon
	// world

	// create a request
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area/", nil)
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

	return nil
}
