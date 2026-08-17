package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func fetchLocationAreas(url string) (LocationAreasAPIResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return LocationAreasAPIResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasAPIResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	var apiResponse LocationAreasAPIResponse

	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return LocationAreasAPIResponse{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return apiResponse, nil
}
