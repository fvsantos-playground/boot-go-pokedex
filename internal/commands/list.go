package commands

import (
	"fmt"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokeapi"
	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokemon"
)

func List(config *pokeapi.Config) error {
	fmt.Println("Your Pokedex:")
	for val := range pokemon.GetNames() {
		fmt.Printf(" - %v\n", val)
	}

	return nil
}
