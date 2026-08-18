package main

import (
	"time"

	"github.com/Lilrosco/pokedexcli/internal/pokecache"
)

func main() {
	cfg := &config {
		commands: getCommands(),
		baseLocationAreasURL: "https://pokeapi.co/api/v2/location-area/",
		cache: pokecache.New(5 * time.Minute),
	}

	startRepl(cfg)
}
