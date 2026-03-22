package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetPokemonInfo(name string) (PokemonInfo, error) {
	locationURL, err := url.JoinPath(PokemonURL, name)
	if err != nil {
		return PokemonInfo{}, err
	}

	val, ok := cache.Get(locationURL)
	if ok {
		return getData[PokemonInfo](val)
	}

	res, err := http.Get(locationURL)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return PokemonInfo{}, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	val, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	cache.Add(locationURL, val)
	return getData[PokemonInfo](val)
}
