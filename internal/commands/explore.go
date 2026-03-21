package commands

import (
	"fmt"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokeapi"
)

func getPokemonNamesIn(area_name string) ([]string, error) {
	locationData, err := pokeapi.GetLocationData(area_name)
	if err != nil {
		return nil, err
	}

	pokeNames := make([]string, 0)
	for _, encounter := range locationData.PokemonEncounters {
		pokeNames = append(pokeNames, encounter.Pokemon.Name)
	}

	return pokeNames, nil
}

func Explore(config *pokeapi.Config) error {
	names, err := getPokemonNamesIn(config.Param)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", config.Param)
	fmt.Println("Found Pokemon:")
	for _, name := range names {
		fmt.Printf(" - %v\n", name)
	}

	return nil
}
