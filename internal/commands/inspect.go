package commands

import (
	"fmt"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokeapi"
	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokemon"
)

func Inspect(config *pokeapi.Config) error {
	pokemon, ok := pokemon.Get(config.Param)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}
