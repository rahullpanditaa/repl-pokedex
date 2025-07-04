package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MapCommand() {
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

	var areas any
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&areas); err != nil {
		log.Fatal(err)
	}

	fmt.Println(areas)

}
