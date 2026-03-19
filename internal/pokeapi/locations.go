package pokeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokecache"
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

var cache pokecache.Cache

func SetCache(c pokecache.Cache) {
	cache = c
}

func GetLocations(locationURL string) (PaginationData, error) {
	getDecodedData := func(val []byte) (PaginationData, error) {
		var data PaginationData
		decoder := json.NewDecoder(bytes.NewReader(val))
		if err := decoder.Decode(&data); err != nil {
			fmt.Println(err)
			return PaginationData{}, err
		}

		return data, nil
	}

	if len(locationURL) == 0 {
		locationURL = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	if val, ok := cache.Get(locationURL); ok {
		return getDecodedData(val)
	}

	res, err := http.Get(locationURL)
	if err != nil {
		fmt.Println(err)
		return PaginationData{}, err
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return PaginationData{}, err
	}

	cache.Add(locationURL, bytes)
	return getDecodedData(bytes)
}
