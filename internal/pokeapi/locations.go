package pokeapi

import (
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
	if len(locationURL) == 0 {
		locationURL = LocationURL + "?offset=0&limit=20"
	}

	if val, ok := cache.Get(locationURL); ok {
		return getData[PaginationData](val)
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
	return getData[PaginationData](bytes)
}
