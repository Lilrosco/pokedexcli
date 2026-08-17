package main

func main() {
	cfg := &config {
		commands: getCommands(),
		baseLocationAreasURL: "https://pokeapi.co/api/v2/location/",
	}

	startRepl(cfg)
}
