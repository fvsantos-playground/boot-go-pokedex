package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PaginationData struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationData `json:"results"`
}

type LocationData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations(locationURL string) (PaginationData, error) {
	if len(locationURL) == 0 {
		locationURL = "https://pokeapi.co/api/v2/location-area"
	}
	res, err := http.Get(locationURL)
	if err != nil {
		fmt.Println(err)
		return PaginationData{}, err
	}
	defer res.Body.Close()

	var data PaginationData
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err)
		return PaginationData{}, err
	}

	return data, nil
}
