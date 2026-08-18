package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Lilrosco/pokedexcli/internal/pokecache"
)

type LocationAreasAPIResponse struct {
	Count int `json:"count`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LocationArea `json:"results"`
}

type LocationArea struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func fetchLocationAreas(url string, cache *pokecache.Cache) (LocationAreasAPIResponse, error) {
	var data []byte

	body, found := cache.Get(url)

	if !found {
		res, err := http.Get(url)

		if err != nil {
			return LocationAreasAPIResponse{}, fmt.Errorf("error creating request: %w", err)
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			return LocationAreasAPIResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		if err != nil {
			return LocationAreasAPIResponse{}, fmt.Errorf("error reading response: %w", err)
		}

		// Save result in cache
		fmt.Println("Caching result")
		cache.Add(url, body)
		data = body
	} else {
		// Cache hit
		fmt.Println("Cache Hit! Loading from memory")
		data = body
	}

	var apiResponse LocationAreasAPIResponse

	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return LocationAreasAPIResponse{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return apiResponse, nil
}
