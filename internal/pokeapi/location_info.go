package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetLocationData(area_name string) (LocationInfoData, error) {
	locationURL, err := url.JoinPath(LocationURL, area_name)
	if err != nil {
		return LocationInfoData{}, err
	}

	val, ok := cache.Get(locationURL)
	if ok {
		return getData[LocationInfoData](val)
	}

	res, err := http.Get(locationURL)
	if err != nil {
		return LocationInfoData{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LocationInfoData{}, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	val, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationInfoData{}, err
	}

	cache.Add(locationURL, val)
	return getData[LocationInfoData](val)
}
