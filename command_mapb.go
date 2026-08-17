package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	url := cfg.baseLocationAreasURL

	if cfg.prevLocationAreasURL == "" {
		fmt.Println("you're on the first page")
	} else {
		url = cfg.prevLocationAreasURL
	}

	apiResponse, err := fetchLocationAreas(url)

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
