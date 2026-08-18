package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	url := cfg.baseLocationAreasURL

	if cfg.nextLocationAreasURL != "" {
		url = cfg.nextLocationAreasURL
	}

	apiResponse, err := fetchLocationAreas(url, cfg.cache)

	if err != nil {
		return err
	}

	// Save Metadata
	cfg.nextLocationAreasURL = apiResponse.Next
	cfg.prevLocationAreasURL = apiResponse.Previous

	for _, locationArea := range apiResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
